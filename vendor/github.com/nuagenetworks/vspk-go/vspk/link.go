/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LinkIdentity represents the Identity of the object
var LinkIdentity = bambou.Identity{
	Name:     "link",
	Category: "links",
}

// LinksList represents a list of Links
type LinksList []*Link

// LinksAncestor is the interface of an ancestor of a Link must implement.
type LinksAncestor interface {
	Links(*bambou.FetchingInfo) (LinksList, *bambou.Error)
	CreateLinks(*Link) *bambou.Error
}

// Link represents the model of a link
type Link struct {
	ID                        string `json:"ID,omitempty"`
	ParentID                  string `json:"parentID,omitempty"`
	ParentType                string `json:"parentType,omitempty"`
	Owner                     string `json:"owner,omitempty"`
	LastUpdatedBy             string `json:"lastUpdatedBy,omitempty"`
	AcceptanceCriteria        string `json:"acceptanceCriteria,omitempty"`
	ReadOnly                  bool   `json:"readOnly"`
	EntityScope               string `json:"entityScope,omitempty"`
	AssociatedDestinationID   string `json:"associatedDestinationID,omitempty"`
	AssociatedDestinationName string `json:"associatedDestinationName,omitempty"`
	AssociatedDestinationType string `json:"associatedDestinationType,omitempty"`
	AssociatedSourceID        string `json:"associatedSourceID,omitempty"`
	AssociatedSourceName      string `json:"associatedSourceName,omitempty"`
	AssociatedSourceType      string `json:"associatedSourceType,omitempty"`
	ExternalID                string `json:"externalID,omitempty"`
	Type                      string `json:"type,omitempty"`
}

// NewLink returns a new *Link
func NewLink() *Link {

	return &Link{}
}

// Identity returns the Identity of the object.
func (o *Link) Identity() bambou.Identity {

	return LinkIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Link) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Link) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Link from the server
func (o *Link) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Link into the server
func (o *Link) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Link from the server
func (o *Link) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Link
func (o *Link) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Link
func (o *Link) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Link
func (o *Link) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Link
func (o *Link) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
