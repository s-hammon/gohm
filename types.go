package main

type ADT struct {
	SendingApp              string `hl7:"MSH.3"`
	SendingFacility         string `hl7:"MSH.4"`
	SendDT                  string `hl7:"MSH.7"`
	MessageType             CM_MSG `hl7:"MSH.9"`
	ControlID               string `hl7:"MSH.10"`
	VersionID               string `hl7:"MSH.12"`
	ExternalPatientID       CX     `hl7:"PID.2"`
	InternalPatientID       CX     `hl7:"PID.3"`
	AlternatePatientID      CX     `hl7:"PID.4"`
	PatientClass            string `hl7:"PV1.2"`
	AssignedPatientLocation PL     `hl7:"PV1.3"`
	ServicingFacility       string `hl7:"PV1.39"`
	MsgPath                 string
}

type ORM struct {
	SendingApp              string `hl7:"MSH.3"`
	SendingFacility         string `hl7:"MSH.4"`
	SendDT                  string `hl7:"MSH.7"`
	MessageType             CM_MSG `hl7:"MSH.9"`
	ControlID               string `hl7:"MSH.10"`
	VersionID               string `hl7:"MSH.12"`
	ExternalPatientID       CX     `hl7:"PID.2"`
	InternalPatientID       CX     `hl7:"PID.3"`
	AlternatePatientID      CX     `hl7:"PID.4"`
	PatientClass            string `hl7:"PV1.2"`
	AssignedPatientLocation PL     `hl7:"PV1.3"`
	ServicingFacility       string `hl7:"PV1.39"`
	OrderControl            string `hl7:"ORC.1"`
	ORCPlacerOrderNumber    string `hl7:"ORC.2"`
	ORCFillerOrderNumber    string `hl7:"ORC.3"`
	OrderStatus             string `hl7:"ORC.5"`
	OBRPlacerOrderNumber    string `hl7:"OBR.2"`
	OBRFillerOrderNumber    string `hl7:"OBR.3"`
	ServiceIdentifier       CE     `hl7:"OBR.4"`
	Priority                string `hl7:"OBR.5"`
	MsgPath                 string
}

type ORU struct {
	SendingApp                 string `hl7:"MSH.3"`
	SendingFacility            string `hl7:"MSH.4"`
	SendDT                     string `hl7:"MSH.7"`
	MessageType                CM_MSG `hl7:"MSH.9"`
	ControlID                  string `hl7:"MSH.10"`
	VersionID                  string `hl7:"MSH.12"`
	ExternalPatientID          CX     `hl7:"PID.2"`
	InternalPatientID          CX     `hl7:"PID.3"`
	AlternatePatientID         CX     `hl7:"PID.4"`
	PatientClass               string `hl7:"PV1.2"`
	AssignedPatientLocation    PL     `hl7:"PV1.3"`
	OrderControl               string `hl7:"ORC.1"`
	ORCPlacerOrderNumber       string `hl7:"ORC.2"`
	ORCFillerOrderNumber       string `hl7:"ORC.3"`
	OrderStatus                string `hl7:"ORC.5"`
	OBRPlacerOrderNumber       string `hl7:"OBR.2"`
	OBRFillerOrderNumber       string `hl7:"OBR.3"`
	ServiceIdentifier          CE     `hl7:"OBR.4"`
	Priority                   string `hl7:"OBR.5"`
	ResultStatus               string `hl7:"OBR.25"`
	PrincipalResultInterpreter CM_NDL `hl7:"OBR.32"`
	MsgPath                    string
}

type MDM struct {
	SendingApp                 string `hl7:"MSH.3"`
	SendingFacility            string `hl7:"MSH.4"`
	SendDT                     string `hl7:"MSH.7"`
	MessageType                CM_MSG `hl7:"MSH.9"`
	ControlID                  string `hl7:"MSH.10"`
	VersionID                  string `hl7:"MSH.12"`
	ExternalPatientID          CX     `hl7:"PID.2"`
	InternalPatientID          CX     `hl7:"PID.3"`
	AlternatePatientID         CX     `hl7:"PID.4"`
	PatientClass               string `hl7:"PV1.2"`
	AssignedPatientLocation    PL     `hl7:"PV1.3"`
	OrderControl               string `hl7:"ORC.1"`
	ORCPlacerOrderNumber       string `hl7:"ORC.2"`
	ORCFillerOrderNumber       string `hl7:"ORC.3"`
	OrderStatus                string `hl7:"ORC.5"`
	OBRPlacerOrderNumber       string `hl7:"OBR.2"`
	OBRFillerOrderNumber       string `hl7:"OBR.3"`
	ServiceIdentifier          CE     `hl7:"OBR.4"`
	Priority                   string `hl7:"OBR.5"`
	ResultStatus               string `hl7:"OBR.25"`
	PrincipalResultInterpreter CM_NDL `hl7:"OBR.32"`
	MsgPath                    string
}

type CE struct {
	Identifier            string `hl7:"1"`
	Text                  string `hl7:"2"`
	CodingSystem          string `hl7:"3"`
	AlternateIdentifier   string `hl7:"4"`
	AlternateText         string `hl7:"5"`
	AlternateCodingSystem string `hl7:"6"`
}

type CN struct {
	IdNumber           string `hl7:"1"`
	FamilyName         string `hl7:"2"`
	GivenName          string `hl7:"3"`
	MiddleName         string `hl7:"4"`
	Suffix             string `hl7:"5"`
	Prefix             string `hl7:"6"`
	Degree             string `hl7:"7"`
	SourceTable        string `hl7:"8"`
	AssigningAuthority string `hl7:"9"`
}

type CX struct {
	ID                 string `hl7:"1"`
	CheckDigit         string `hl7:"2"`
	CheckDigitScheme   string `hl7:"3"`
	AssigningAuthority string `hl7:"4"`
	IdentifierTypeCode string `hl7:"5"`
}

type HD struct {
	NamespaceID     string `hl7:"1"`
	UniversalID     string `hl7:"2"`
	UniversalIDType string `hl7:"3"`
}

type PL struct {
	PointOfCare         string `hl7:"1"`
	Room                string `hl7:"2"`
	Bed                 string `hl7:"3"`
	Facility            string `hl7:"4"`
	LocationStatus      string `hl7:"5"`
	PersonLocationType  string `hl7:"6"`
	Building            string `hl7:"7"`
	Floor               string `hl7:"8"`
	LocationDescription string `hl7:"9"`
}

type CM_MSG struct {
	Name         string `hl7:"1"`
	TriggerEvent string `hl7:"2"`
}

type CM_NDL struct {
	OpName             CN     `hl7:"1"`
	StartDT            string `hl7:"2"`
	EndDT              string `hl7:"3"`
	PointOfCare        string `hl7:"4"`
	Room               string `hl7:"5"`
	Bed                string `hl7:"6"`
	Facility           string `hl7:"7"`
	LocationStatus     string `hl7:"8"`
	PersonLocationType string `hl7:"9"`
	Building           string `hl7:"10"`
	Floor              string `hl7:"11"`
}
