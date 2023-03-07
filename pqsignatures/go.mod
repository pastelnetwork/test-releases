module github.com/pastelnetwork/gonode/pqsignatures

go 1.16

require (
	github.com/DataDog/zstd v1.4.8
	github.com/fogleman/gg v1.3.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/kevinburke/nacl v0.0.0-20210321052800-030051251ea5
	github.com/makiuchi-d/gozxing v0.0.0-20210324052758-57132e828831
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/pastelnetwork/gonode/common v0.0.0
	github.com/pastelnetwork/gonode/legroast v0.0.0
	github.com/pastelnetwork/steganography v1.0.1
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/xlzd/gotp v0.0.0-20181030022105-c8557ba2c119
	golang.org/x/crypto v0.1.0
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
)

replace github.com/pastelnetwork/gonode/common => ../common

replace github.com/pastelnetwork/gonode/legroast => ../legroast
