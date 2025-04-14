package wire_pkg_demo

import (
	"context"
	"errors"
	"fmt"
)

type C struct {
	Content string
}

func NewC(ctx context.Context, b B) (C, error) {
	if b.Desc == "" {
		return C{}, errors.New("描述信息不能为空哦")
	}

	return C{Content: fmt.Sprintf("现在的描述信息是: %s", b.Desc)}, nil
}
