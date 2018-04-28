package bddraft

import(
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestServerRun(t *testing.T) {
	Convey("",t, func() {
		So(1+1,ShouldEqual,2)
	})
}

func TestClient(t *testing.T) {
	Convey("客户端",t, func() {

		Convey("添加一个键值对", func() {
			So("ok!",ShouldContainSubstring,"ok!")
		})

		Convey("获取一个键",func(){
			So(1+1,ShouldEqual,2)
		})

	})
}



func TestServer(t *testing.T)  {
	Convey("服务端",t,func(){
		Convey("启动服务", func() {
			Convey("添加处理器", func() {

				Convey("添加Add处理器", func() {

				})

				Convey("添加Get处理器", func() {

				})
			})
		})
	})

}


