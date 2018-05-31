// The person Type (more like an object)
package main

type Address1 struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}