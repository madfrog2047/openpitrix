// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package repoiface

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"

	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

type HttpInterface struct {
	url *neturl.URL
}

func NewHttpInterface(ctx context.Context, url *neturl.URL) (*HttpInterface, error) {
	return &HttpInterface{
		url: url,
	}, nil
}

func (i *HttpInterface) ReadFile(ctx context.Context, filename string) ([]byte, error) {
	u := strings.TrimSuffix(i.url.String(), "/") + "/" + filename

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status code is not 200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (i *HttpInterface) WriteFile(ctx context.Context, filename string, data []byte) error {
	return ErrWriteIsUnsupported
}

type indexYaml struct {
	ApiVersion string                 `yaml:"apiVersion"`
	Entries    map[string]interface{} `yaml:"entries"`
	Generated  string                 `yaml:"generated"`
}

func (i *HttpInterface) CheckRead(ctx context.Context) error {
	data, err := i.ReadFile(ctx, "index.yaml")
	if err != nil {
		return err
	}
	var y indexYaml
	err = yamlutil.Decode(data, &y)
	return err
}

func (i *HttpInterface) CheckWrite(ctx context.Context) error {
	return ErrWriteIsUnsupported
}
