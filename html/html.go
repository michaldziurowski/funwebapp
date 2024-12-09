package html

import "fmt"

type Elem interface {
	String() string
}

func renderElem(tagName string, attrs map[string]string, children []Elem) string {
	attr := ""
	for n, v := range attrs {
		attr += fmt.Sprintf(` %s=%q`, n, v)
	}
	childrenHtml := ""
	for _, child := range children {
		childrenHtml += child.String()
	}

	if childrenHtml == "" {
		return fmt.Sprintf("<%s%s />", tagName, attr)
	}
	return fmt.Sprintf("<%[1]s%s>%s</%[1]s>", tagName, attr, childrenHtml)
}

type Text string

func (x Text) String() string {
	return string(x)
}

type a struct {
	Attrs    map[string]string
	Children []Elem
}

var A *a

func (x *a) String() string {
	return renderElem("a", x.Attrs, x.Children)
}

func (x *a) C(e ...Elem) *a {
	z := x
	if x == nil {
		z = &a{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *a) Href(v string) *a {
	z := x
	if x == nil {
		z = &a{Attrs: map[string]string{}}
	}
	z.Attrs["href"] = v
	return z
}
func (x *a) Id(v string) *a {
	z := x
	if x == nil {
		z = &a{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *a) Class(v string) *a {
	z := x
	if x == nil {
		z = &a{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type body struct {
	Attrs    map[string]string
	Children []Elem
}

var Body *body

func (x *body) String() string {
	return renderElem("body", x.Attrs, x.Children)
}

func (x *body) C(e ...Elem) *body {
	z := x
	if x == nil {
		z = &body{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *body) Id(v string) *body {
	z := x
	if x == nil {
		z = &body{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *body) Class(v string) *body {
	z := x
	if x == nil {
		z = &body{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type button struct {
	Attrs    map[string]string
	Children []Elem
}

var Button *button

func (x *button) String() string {
	return renderElem("button", x.Attrs, x.Children)
}

func (x *button) C(e ...Elem) *button {
	z := x
	if x == nil {
		z = &button{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *button) Type(v string) *button {
	z := x
	if x == nil {
		z = &button{Attrs: map[string]string{}}
	}
	z.Attrs["type"] = v
	return z
}
func (x *button) Id(v string) *button {
	z := x
	if x == nil {
		z = &button{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *button) Class(v string) *button {
	z := x
	if x == nil {
		z = &button{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type div struct {
	Attrs    map[string]string
	Children []Elem
}

var Div *div

func (x *div) String() string {
	return renderElem("div", x.Attrs, x.Children)
}

func (x *div) C(e ...Elem) *div {
	z := x
	if x == nil {
		z = &div{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *div) Id(v string) *div {
	z := x
	if x == nil {
		z = &div{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *div) Class(v string) *div {
	z := x
	if x == nil {
		z = &div{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type form struct {
	Attrs    map[string]string
	Children []Elem
}

var Form *form

func (x *form) String() string {
	return renderElem("form", x.Attrs, x.Children)
}

func (x *form) C(e ...Elem) *form {
	z := x
	if x == nil {
		z = &form{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *form) Action(v string) *form {
	z := x
	if x == nil {
		z = &form{Attrs: map[string]string{}}
	}
	z.Attrs["action"] = v
	return z
}
func (x *form) Id(v string) *form {
	z := x
	if x == nil {
		z = &form{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *form) Class(v string) *form {
	z := x
	if x == nil {
		z = &form{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type head struct {
	Attrs    map[string]string
	Children []Elem
}

var Head *head

func (x *head) String() string {
	return renderElem("head", x.Attrs, x.Children)
}

func (x *head) C(e ...Elem) *head {
	z := x
	if x == nil {
		z = &head{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *head) Id(v string) *head {
	z := x
	if x == nil {
		z = &head{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *head) Class(v string) *head {
	z := x
	if x == nil {
		z = &head{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type html struct {
	Attrs    map[string]string
	Children []Elem
}

var Html *html

func (x *html) String() string {
	return renderElem("html", x.Attrs, x.Children)
}

func (x *html) C(e ...Elem) *html {
	z := x
	if x == nil {
		z = &html{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *html) Lang(v string) *html {
	z := x
	if x == nil {
		z = &html{Attrs: map[string]string{}}
	}
	z.Attrs["lang"] = v
	return z
}
func (x *html) Id(v string) *html {
	z := x
	if x == nil {
		z = &html{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *html) Class(v string) *html {
	z := x
	if x == nil {
		z = &html{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type img struct {
	Attrs    map[string]string
	Children []Elem
}

var Img *img

func (x *img) String() string {
	return renderElem("img", x.Attrs, x.Children)
}

func (x *img) C(e ...Elem) *img {
	z := x
	if x == nil {
		z = &img{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *img) Src(v string) *img {
	z := x
	if x == nil {
		z = &img{Attrs: map[string]string{}}
	}
	z.Attrs["src"] = v
	return z
}
func (x *img) Alt(v string) *img {
	z := x
	if x == nil {
		z = &img{Attrs: map[string]string{}}
	}
	z.Attrs["alt"] = v
	return z
}
func (x *img) Id(v string) *img {
	z := x
	if x == nil {
		z = &img{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *img) Class(v string) *img {
	z := x
	if x == nil {
		z = &img{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type input struct {
	Attrs    map[string]string
	Children []Elem
}

var Input *input

func (x *input) String() string {
	return renderElem("input", x.Attrs, x.Children)
}

func (x *input) C(e ...Elem) *input {
	z := x
	if x == nil {
		z = &input{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *input) Type(v string) *input {
	z := x
	if x == nil {
		z = &input{Attrs: map[string]string{}}
	}
	z.Attrs["type"] = v
	return z
}
func (x *input) Value(v string) *input {
	z := x
	if x == nil {
		z = &input{Attrs: map[string]string{}}
	}
	z.Attrs["value"] = v
	return z
}
func (x *input) Id(v string) *input {
	z := x
	if x == nil {
		z = &input{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *input) Class(v string) *input {
	z := x
	if x == nil {
		z = &input{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type li struct {
	Attrs    map[string]string
	Children []Elem
}

var Li *li

func (x *li) String() string {
	return renderElem("li", x.Attrs, x.Children)
}

func (x *li) C(e ...Elem) *li {
	z := x
	if x == nil {
		z = &li{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *li) Id(v string) *li {
	z := x
	if x == nil {
		z = &li{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *li) Class(v string) *li {
	z := x
	if x == nil {
		z = &li{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type link struct {
	Attrs    map[string]string
	Children []Elem
}

var Link *link

func (x *link) String() string {
	return renderElem("link", x.Attrs, x.Children)
}

func (x *link) C(e ...Elem) *link {
	z := x
	if x == nil {
		z = &link{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *link) Rel(v string) *link {
	z := x
	if x == nil {
		z = &link{Attrs: map[string]string{}}
	}
	z.Attrs["rel"] = v
	return z
}
func (x *link) Href(v string) *link {
	z := x
	if x == nil {
		z = &link{Attrs: map[string]string{}}
	}
	z.Attrs["href"] = v
	return z
}
func (x *link) Id(v string) *link {
	z := x
	if x == nil {
		z = &link{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *link) Class(v string) *link {
	z := x
	if x == nil {
		z = &link{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type main struct {
	Attrs    map[string]string
	Children []Elem
}

var Main *main

func (x *main) String() string {
	return renderElem("main", x.Attrs, x.Children)
}

func (x *main) C(e ...Elem) *main {
	z := x
	if x == nil {
		z = &main{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *main) Id(v string) *main {
	z := x
	if x == nil {
		z = &main{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *main) Class(v string) *main {
	z := x
	if x == nil {
		z = &main{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type meta struct {
	Attrs    map[string]string
	Children []Elem
}

var Meta *meta

func (x *meta) String() string {
	return renderElem("meta", x.Attrs, x.Children)
}

func (x *meta) C(e ...Elem) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *meta) Name(v string) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Attrs["name"] = v
	return z
}
func (x *meta) Content(v string) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Attrs["content"] = v
	return z
}
func (x *meta) Charset(v string) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Attrs["charset"] = v
	return z
}
func (x *meta) Id(v string) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *meta) Class(v string) *meta {
	z := x
	if x == nil {
		z = &meta{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type p struct {
	Attrs    map[string]string
	Children []Elem
}

var P *p

func (x *p) String() string {
	return renderElem("p", x.Attrs, x.Children)
}

func (x *p) C(e ...Elem) *p {
	z := x
	if x == nil {
		z = &p{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *p) Id(v string) *p {
	z := x
	if x == nil {
		z = &p{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *p) Class(v string) *p {
	z := x
	if x == nil {
		z = &p{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type script struct {
	Attrs    map[string]string
	Children []Elem
}

var Script *script

func (x *script) String() string {
	return renderElem("script", x.Attrs, x.Children)
}

func (x *script) C(e ...Elem) *script {
	z := x
	if x == nil {
		z = &script{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *script) Src(v string) *script {
	z := x
	if x == nil {
		z = &script{Attrs: map[string]string{}}
	}
	z.Attrs["src"] = v
	return z
}
func (x *script) Id(v string) *script {
	z := x
	if x == nil {
		z = &script{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *script) Class(v string) *script {
	z := x
	if x == nil {
		z = &script{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type span struct {
	Attrs    map[string]string
	Children []Elem
}

var Span *span

func (x *span) String() string {
	return renderElem("span", x.Attrs, x.Children)
}

func (x *span) C(e ...Elem) *span {
	z := x
	if x == nil {
		z = &span{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *span) Id(v string) *span {
	z := x
	if x == nil {
		z = &span{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *span) Class(v string) *span {
	z := x
	if x == nil {
		z = &span{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type style struct {
	Attrs    map[string]string
	Children []Elem
}

var Style *style

func (x *style) String() string {
	return renderElem("style", x.Attrs, x.Children)
}

func (x *style) C(e ...Elem) *style {
	z := x
	if x == nil {
		z = &style{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *style) Id(v string) *style {
	z := x
	if x == nil {
		z = &style{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *style) Class(v string) *style {
	z := x
	if x == nil {
		z = &style{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type title struct {
	Attrs    map[string]string
	Children []Elem
}

var Title *title

func (x *title) String() string {
	return renderElem("title", x.Attrs, x.Children)
}

func (x *title) C(e ...Elem) *title {
	z := x
	if x == nil {
		z = &title{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *title) Id(v string) *title {
	z := x
	if x == nil {
		z = &title{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *title) Class(v string) *title {
	z := x
	if x == nil {
		z = &title{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}

type ul struct {
	Attrs    map[string]string
	Children []Elem
}

var Ul *ul

func (x *ul) String() string {
	return renderElem("ul", x.Attrs, x.Children)
}

func (x *ul) C(e ...Elem) *ul {
	z := x
	if x == nil {
		z = &ul{Attrs: map[string]string{}}
	}
	z.Children = append(z.Children, e...)
	return z
}

func (x *ul) Id(v string) *ul {
	z := x
	if x == nil {
		z = &ul{Attrs: map[string]string{}}
	}
	z.Attrs["id"] = v
	return z
}
func (x *ul) Class(v string) *ul {
	z := x
	if x == nil {
		z = &ul{Attrs: map[string]string{}}
	}
	z.Attrs["class"] = v
	return z
}
