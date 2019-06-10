// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Discovery": Application Media Types
//
// Command:
// $ goagen
// --design=ontt/design
// --out=$(GOPATH)/src/ontt
// --version=v1.3.1

package client

import (
	"net/http"
	"time"
)

// NeatThing media type (default view)
//
// Identifier: application/vnd.douthitlab.neatthing; view=default
type NeatThing struct {
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" yaml:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" yaml:"link,omitempty" xml:"link,omitempty"`
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// NeatThing media type (full view)
//
// Identifier: application/vnd.douthitlab.neatthing; view=full
type NeatThingFull struct {
	Bibliography []string `form:"bibliography,omitempty" json:"bibliography,omitempty" yaml:"bibliography,omitempty" xml:"bibliography,omitempty"`
	// When this was a neat thing
	Date *time.Time `form:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty" xml:"date,omitempty"`
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" yaml:"definition,omitempty" xml:"definition,omitempty"`
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" yaml:"link,omitempty" xml:"link,omitempty"`
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// NeatThing media type (name view)
//
// Identifier: application/vnd.douthitlab.neatthing; view=name
type NeatThingName struct {
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// NeatThing media type (name+definition view)
//
// Identifier: application/vnd.douthitlab.neatthing; view=name+definition
type NeatThingNameDefinition struct {
	// What the neat thing is
	Definition *string `form:"definition,omitempty" json:"definition,omitempty" yaml:"definition,omitempty" xml:"definition,omitempty"`
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// NeatThing media type (name+link view)
//
// Identifier: application/vnd.douthitlab.neatthing; view=name+link
type NeatThingNameLink struct {
	// Illustrative link for the neat thing
	Link *string `form:"link,omitempty" json:"link,omitempty" yaml:"link,omitempty" xml:"link,omitempty"`
	// The neat thing
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// DecodeNeatThing decodes the NeatThing instance encoded in resp body.
func (c *Client) DecodeNeatThing(resp *http.Response) (*NeatThing, error) {
	var decoded NeatThing
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeNeatThingFull decodes the NeatThingFull instance encoded in resp body.
func (c *Client) DecodeNeatThingFull(resp *http.Response) (*NeatThingFull, error) {
	var decoded NeatThingFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeNeatThingName decodes the NeatThingName instance encoded in resp body.
func (c *Client) DecodeNeatThingName(resp *http.Response) (*NeatThingName, error) {
	var decoded NeatThingName
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeNeatThingNameDefinition decodes the NeatThingNameDefinition instance encoded in resp body.
func (c *Client) DecodeNeatThingNameDefinition(resp *http.Response) (*NeatThingNameDefinition, error) {
	var decoded NeatThingNameDefinition
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeNeatThingNameLink decodes the NeatThingNameLink instance encoded in resp body.
func (c *Client) DecodeNeatThingNameLink(resp *http.Response) (*NeatThingNameLink, error) {
	var decoded NeatThingNameLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
