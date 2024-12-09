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

type elema struct {
	attrs    map[string]string
	children []Elem
}

var A *elema

func (x *elema) String() string {
	return renderElem("a", x.attrs, x.children)
}

func (x *elema) C(e ...Elem) *elema {
	z := x
	if x == nil {
		z = &elema{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elema) Href(v string) *elema {
	z := x
	if x == nil {
		z = &elema{attrs: map[string]string{}}
	}
	z.attrs["href"] = v
	return z
}
func (x *elema) Id(v string) *elema {
	z := x
	if x == nil {
		z = &elema{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elema) Class(v string) *elema {
	z := x
	if x == nil {
		z = &elema{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemarticle struct {
	attrs    map[string]string
	children []Elem
}

var Article *elemarticle

func (x *elemarticle) String() string {
	return renderElem("article", x.attrs, x.children)
}

func (x *elemarticle) C(e ...Elem) *elemarticle {
	z := x
	if x == nil {
		z = &elemarticle{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemarticle) Id(v string) *elemarticle {
	z := x
	if x == nil {
		z = &elemarticle{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemarticle) Class(v string) *elemarticle {
	z := x
	if x == nil {
		z = &elemarticle{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elembody struct {
	attrs    map[string]string
	children []Elem
}

var Body *elembody

func (x *elembody) String() string {
	return renderElem("body", x.attrs, x.children)
}

func (x *elembody) C(e ...Elem) *elembody {
	z := x
	if x == nil {
		z = &elembody{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elembody) Id(v string) *elembody {
	z := x
	if x == nil {
		z = &elembody{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elembody) Class(v string) *elembody {
	z := x
	if x == nil {
		z = &elembody{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elembutton struct {
	attrs    map[string]string
	children []Elem
}

var Button *elembutton

func (x *elembutton) String() string {
	return renderElem("button", x.attrs, x.children)
}

func (x *elembutton) C(e ...Elem) *elembutton {
	z := x
	if x == nil {
		z = &elembutton{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elembutton) Type(v string) *elembutton {
	z := x
	if x == nil {
		z = &elembutton{attrs: map[string]string{}}
	}
	z.attrs["type"] = v
	return z
}
func (x *elembutton) Id(v string) *elembutton {
	z := x
	if x == nil {
		z = &elembutton{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elembutton) Class(v string) *elembutton {
	z := x
	if x == nil {
		z = &elembutton{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemdetails struct {
	attrs    map[string]string
	children []Elem
}

var Details *elemdetails

func (x *elemdetails) String() string {
	return renderElem("details", x.attrs, x.children)
}

func (x *elemdetails) C(e ...Elem) *elemdetails {
	z := x
	if x == nil {
		z = &elemdetails{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemdetails) Id(v string) *elemdetails {
	z := x
	if x == nil {
		z = &elemdetails{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemdetails) Class(v string) *elemdetails {
	z := x
	if x == nil {
		z = &elemdetails{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemdiv struct {
	attrs    map[string]string
	children []Elem
}

var Div *elemdiv

func (x *elemdiv) String() string {
	return renderElem("div", x.attrs, x.children)
}

func (x *elemdiv) C(e ...Elem) *elemdiv {
	z := x
	if x == nil {
		z = &elemdiv{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemdiv) Id(v string) *elemdiv {
	z := x
	if x == nil {
		z = &elemdiv{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemdiv) Class(v string) *elemdiv {
	z := x
	if x == nil {
		z = &elemdiv{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemfooter struct {
	attrs    map[string]string
	children []Elem
}

var Footer *elemfooter

func (x *elemfooter) String() string {
	return renderElem("footer", x.attrs, x.children)
}

func (x *elemfooter) C(e ...Elem) *elemfooter {
	z := x
	if x == nil {
		z = &elemfooter{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemfooter) Id(v string) *elemfooter {
	z := x
	if x == nil {
		z = &elemfooter{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemfooter) Class(v string) *elemfooter {
	z := x
	if x == nil {
		z = &elemfooter{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemform struct {
	attrs    map[string]string
	children []Elem
}

var Form *elemform

func (x *elemform) String() string {
	return renderElem("form", x.attrs, x.children)
}

func (x *elemform) C(e ...Elem) *elemform {
	z := x
	if x == nil {
		z = &elemform{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemform) Action(v string) *elemform {
	z := x
	if x == nil {
		z = &elemform{attrs: map[string]string{}}
	}
	z.attrs["action"] = v
	return z
}
func (x *elemform) Id(v string) *elemform {
	z := x
	if x == nil {
		z = &elemform{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemform) Class(v string) *elemform {
	z := x
	if x == nil {
		z = &elemform{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemhead struct {
	attrs    map[string]string
	children []Elem
}

var Head *elemhead

func (x *elemhead) String() string {
	return renderElem("head", x.attrs, x.children)
}

func (x *elemhead) C(e ...Elem) *elemhead {
	z := x
	if x == nil {
		z = &elemhead{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemhead) Id(v string) *elemhead {
	z := x
	if x == nil {
		z = &elemhead{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemhead) Class(v string) *elemhead {
	z := x
	if x == nil {
		z = &elemhead{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemheader struct {
	attrs    map[string]string
	children []Elem
}

var Header *elemheader

func (x *elemheader) String() string {
	return renderElem("header", x.attrs, x.children)
}

func (x *elemheader) C(e ...Elem) *elemheader {
	z := x
	if x == nil {
		z = &elemheader{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemheader) Id(v string) *elemheader {
	z := x
	if x == nil {
		z = &elemheader{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemheader) Class(v string) *elemheader {
	z := x
	if x == nil {
		z = &elemheader{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemhtml struct {
	attrs    map[string]string
	children []Elem
}

var Html *elemhtml

func (x *elemhtml) String() string {
	return renderElem("html", x.attrs, x.children)
}

func (x *elemhtml) C(e ...Elem) *elemhtml {
	z := x
	if x == nil {
		z = &elemhtml{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemhtml) Lang(v string) *elemhtml {
	z := x
	if x == nil {
		z = &elemhtml{attrs: map[string]string{}}
	}
	z.attrs["lang"] = v
	return z
}
func (x *elemhtml) Id(v string) *elemhtml {
	z := x
	if x == nil {
		z = &elemhtml{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemhtml) Class(v string) *elemhtml {
	z := x
	if x == nil {
		z = &elemhtml{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemimg struct {
	attrs    map[string]string
	children []Elem
}

var Img *elemimg

func (x *elemimg) String() string {
	return renderElem("img", x.attrs, x.children)
}

func (x *elemimg) C(e ...Elem) *elemimg {
	z := x
	if x == nil {
		z = &elemimg{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemimg) Src(v string) *elemimg {
	z := x
	if x == nil {
		z = &elemimg{attrs: map[string]string{}}
	}
	z.attrs["src"] = v
	return z
}
func (x *elemimg) Alt(v string) *elemimg {
	z := x
	if x == nil {
		z = &elemimg{attrs: map[string]string{}}
	}
	z.attrs["alt"] = v
	return z
}
func (x *elemimg) Id(v string) *elemimg {
	z := x
	if x == nil {
		z = &elemimg{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemimg) Class(v string) *elemimg {
	z := x
	if x == nil {
		z = &elemimg{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type eleminput struct {
	attrs    map[string]string
	children []Elem
}

var Input *eleminput

func (x *eleminput) String() string {
	return renderElem("input", x.attrs, x.children)
}

func (x *eleminput) C(e ...Elem) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *eleminput) Type(v string) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.attrs["type"] = v
	return z
}
func (x *eleminput) Value(v string) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.attrs["value"] = v
	return z
}
func (x *eleminput) Name(v string) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.attrs["name"] = v
	return z
}
func (x *eleminput) Id(v string) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *eleminput) Class(v string) *eleminput {
	z := x
	if x == nil {
		z = &eleminput{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemlabel struct {
	attrs    map[string]string
	children []Elem
}

var Label *elemlabel

func (x *elemlabel) String() string {
	return renderElem("label", x.attrs, x.children)
}

func (x *elemlabel) C(e ...Elem) *elemlabel {
	z := x
	if x == nil {
		z = &elemlabel{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemlabel) For(v string) *elemlabel {
	z := x
	if x == nil {
		z = &elemlabel{attrs: map[string]string{}}
	}
	z.attrs["for"] = v
	return z
}
func (x *elemlabel) Id(v string) *elemlabel {
	z := x
	if x == nil {
		z = &elemlabel{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemlabel) Class(v string) *elemlabel {
	z := x
	if x == nil {
		z = &elemlabel{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemli struct {
	attrs    map[string]string
	children []Elem
}

var Li *elemli

func (x *elemli) String() string {
	return renderElem("li", x.attrs, x.children)
}

func (x *elemli) C(e ...Elem) *elemli {
	z := x
	if x == nil {
		z = &elemli{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemli) Id(v string) *elemli {
	z := x
	if x == nil {
		z = &elemli{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemli) Class(v string) *elemli {
	z := x
	if x == nil {
		z = &elemli{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemlink struct {
	attrs    map[string]string
	children []Elem
}

var Link *elemlink

func (x *elemlink) String() string {
	return renderElem("link", x.attrs, x.children)
}

func (x *elemlink) C(e ...Elem) *elemlink {
	z := x
	if x == nil {
		z = &elemlink{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemlink) Rel(v string) *elemlink {
	z := x
	if x == nil {
		z = &elemlink{attrs: map[string]string{}}
	}
	z.attrs["rel"] = v
	return z
}
func (x *elemlink) Href(v string) *elemlink {
	z := x
	if x == nil {
		z = &elemlink{attrs: map[string]string{}}
	}
	z.attrs["href"] = v
	return z
}
func (x *elemlink) Id(v string) *elemlink {
	z := x
	if x == nil {
		z = &elemlink{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemlink) Class(v string) *elemlink {
	z := x
	if x == nil {
		z = &elemlink{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemmain struct {
	attrs    map[string]string
	children []Elem
}

var Main *elemmain

func (x *elemmain) String() string {
	return renderElem("main", x.attrs, x.children)
}

func (x *elemmain) C(e ...Elem) *elemmain {
	z := x
	if x == nil {
		z = &elemmain{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemmain) Id(v string) *elemmain {
	z := x
	if x == nil {
		z = &elemmain{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemmain) Class(v string) *elemmain {
	z := x
	if x == nil {
		z = &elemmain{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemmeta struct {
	attrs    map[string]string
	children []Elem
}

var Meta *elemmeta

func (x *elemmeta) String() string {
	return renderElem("meta", x.attrs, x.children)
}

func (x *elemmeta) C(e ...Elem) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemmeta) Name(v string) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.attrs["name"] = v
	return z
}
func (x *elemmeta) Content(v string) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.attrs["content"] = v
	return z
}
func (x *elemmeta) Charset(v string) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.attrs["charset"] = v
	return z
}
func (x *elemmeta) Id(v string) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemmeta) Class(v string) *elemmeta {
	z := x
	if x == nil {
		z = &elemmeta{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemnav struct {
	attrs    map[string]string
	children []Elem
}

var Nav *elemnav

func (x *elemnav) String() string {
	return renderElem("nav", x.attrs, x.children)
}

func (x *elemnav) C(e ...Elem) *elemnav {
	z := x
	if x == nil {
		z = &elemnav{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemnav) Id(v string) *elemnav {
	z := x
	if x == nil {
		z = &elemnav{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemnav) Class(v string) *elemnav {
	z := x
	if x == nil {
		z = &elemnav{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemoption struct {
	attrs    map[string]string
	children []Elem
}

var Option *elemoption

func (x *elemoption) String() string {
	return renderElem("option", x.attrs, x.children)
}

func (x *elemoption) C(e ...Elem) *elemoption {
	z := x
	if x == nil {
		z = &elemoption{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemoption) Id(v string) *elemoption {
	z := x
	if x == nil {
		z = &elemoption{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemoption) Class(v string) *elemoption {
	z := x
	if x == nil {
		z = &elemoption{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemp struct {
	attrs    map[string]string
	children []Elem
}

var P *elemp

func (x *elemp) String() string {
	return renderElem("p", x.attrs, x.children)
}

func (x *elemp) C(e ...Elem) *elemp {
	z := x
	if x == nil {
		z = &elemp{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemp) Id(v string) *elemp {
	z := x
	if x == nil {
		z = &elemp{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemp) Class(v string) *elemp {
	z := x
	if x == nil {
		z = &elemp{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemscript struct {
	attrs    map[string]string
	children []Elem
}

var Script *elemscript

func (x *elemscript) String() string {
	return renderElem("script", x.attrs, x.children)
}

func (x *elemscript) C(e ...Elem) *elemscript {
	z := x
	if x == nil {
		z = &elemscript{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemscript) Src(v string) *elemscript {
	z := x
	if x == nil {
		z = &elemscript{attrs: map[string]string{}}
	}
	z.attrs["src"] = v
	return z
}
func (x *elemscript) Id(v string) *elemscript {
	z := x
	if x == nil {
		z = &elemscript{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemscript) Class(v string) *elemscript {
	z := x
	if x == nil {
		z = &elemscript{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemsection struct {
	attrs    map[string]string
	children []Elem
}

var Section *elemsection

func (x *elemsection) String() string {
	return renderElem("section", x.attrs, x.children)
}

func (x *elemsection) C(e ...Elem) *elemsection {
	z := x
	if x == nil {
		z = &elemsection{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemsection) Id(v string) *elemsection {
	z := x
	if x == nil {
		z = &elemsection{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemsection) Class(v string) *elemsection {
	z := x
	if x == nil {
		z = &elemsection{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemselect struct {
	attrs    map[string]string
	children []Elem
}

var Select *elemselect

func (x *elemselect) String() string {
	return renderElem("select", x.attrs, x.children)
}

func (x *elemselect) C(e ...Elem) *elemselect {
	z := x
	if x == nil {
		z = &elemselect{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemselect) Name(v string) *elemselect {
	z := x
	if x == nil {
		z = &elemselect{attrs: map[string]string{}}
	}
	z.attrs["name"] = v
	return z
}
func (x *elemselect) Id(v string) *elemselect {
	z := x
	if x == nil {
		z = &elemselect{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemselect) Class(v string) *elemselect {
	z := x
	if x == nil {
		z = &elemselect{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemspan struct {
	attrs    map[string]string
	children []Elem
}

var Span *elemspan

func (x *elemspan) String() string {
	return renderElem("span", x.attrs, x.children)
}

func (x *elemspan) C(e ...Elem) *elemspan {
	z := x
	if x == nil {
		z = &elemspan{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemspan) Id(v string) *elemspan {
	z := x
	if x == nil {
		z = &elemspan{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemspan) Class(v string) *elemspan {
	z := x
	if x == nil {
		z = &elemspan{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemstyle struct {
	attrs    map[string]string
	children []Elem
}

var Style *elemstyle

func (x *elemstyle) String() string {
	return renderElem("style", x.attrs, x.children)
}

func (x *elemstyle) C(e ...Elem) *elemstyle {
	z := x
	if x == nil {
		z = &elemstyle{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemstyle) Id(v string) *elemstyle {
	z := x
	if x == nil {
		z = &elemstyle{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemstyle) Class(v string) *elemstyle {
	z := x
	if x == nil {
		z = &elemstyle{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemsummary struct {
	attrs    map[string]string
	children []Elem
}

var Summary *elemsummary

func (x *elemsummary) String() string {
	return renderElem("summary", x.attrs, x.children)
}

func (x *elemsummary) C(e ...Elem) *elemsummary {
	z := x
	if x == nil {
		z = &elemsummary{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemsummary) Id(v string) *elemsummary {
	z := x
	if x == nil {
		z = &elemsummary{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemsummary) Class(v string) *elemsummary {
	z := x
	if x == nil {
		z = &elemsummary{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemtitle struct {
	attrs    map[string]string
	children []Elem
}

var Title *elemtitle

func (x *elemtitle) String() string {
	return renderElem("title", x.attrs, x.children)
}

func (x *elemtitle) C(e ...Elem) *elemtitle {
	z := x
	if x == nil {
		z = &elemtitle{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemtitle) Id(v string) *elemtitle {
	z := x
	if x == nil {
		z = &elemtitle{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemtitle) Class(v string) *elemtitle {
	z := x
	if x == nil {
		z = &elemtitle{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}

type elemul struct {
	attrs    map[string]string
	children []Elem
}

var Ul *elemul

func (x *elemul) String() string {
	return renderElem("ul", x.attrs, x.children)
}

func (x *elemul) C(e ...Elem) *elemul {
	z := x
	if x == nil {
		z = &elemul{attrs: map[string]string{}}
	}
	z.children = append(z.children, e...)
	return z
}

func (x *elemul) Id(v string) *elemul {
	z := x
	if x == nil {
		z = &elemul{attrs: map[string]string{}}
	}
	z.attrs["id"] = v
	return z
}
func (x *elemul) Class(v string) *elemul {
	z := x
	if x == nil {
		z = &elemul{attrs: map[string]string{}}
	}
	z.attrs["class"] = v
	return z
}
