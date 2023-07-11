//go:generate mockgen -source ${GOFILE} -destination xxx_mocks_test.go -package ${GOPACKAGE}_test

package badges

type handler interface {
	IsEnabled() bool
}
