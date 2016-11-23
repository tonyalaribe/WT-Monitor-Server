package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Register is a  helper function that registers a handler under a group. This helper by default registers all the CRUD methods (GET, POST, PUT, etc) by default, and returns and appropriate error message when they are called. Except those default error handlers are replaced by a handler which satisfies the Resource interface, by having a method that overwrites the corresponding default method on the interface.
func Register(p string, r Resource, g *gin.RouterGroup) {
	g.GET(p, r.Get)
	g.POST(p, r.Post)
	g.HEAD(p, r.Head)
	g.PUT(p, r.Put)
	g.PATCH(p, r.Patch)
	g.DELETE(p, r.Delete)
	g.OPTIONS(p, r.Options)
}

//Resource handler helps give all the methods default error messages unless overwritten
type Resource interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Head(c *gin.Context)
	Delete(c *gin.Context)
	Patch(c *gin.Context)
	Put(c *gin.Context)
	Options(c *gin.Context)
}

type Base struct {
	Resource
}

type Message struct {
	APIMessage string `json:"apiMessage"`
}

func (Base) Get(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Post(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Head(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Delete(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Put(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Patch(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}

func (Base) Options(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Message{APIMessage: "not allowed"})
}
