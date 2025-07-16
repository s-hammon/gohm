package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/bigquery"
	"github.com/s-hammon/p"
	"google.golang.org/api/healthcare/v1"
)

type Hl7Service struct {
	client     *bigquery.Client
	msgService *healthcare.ProjectsLocationsDatasetsHl7V2StoresMessagesService
}

func NewHl7Service(client *bigquery.Client, msgService *healthcare.ProjectsLocationsDatasetsHl7V2StoresMessagesService) *Hl7Service {
	return &Hl7Service{client, msgService}
}

func (s *Hl7Service) Mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", s.handleMessage)

	return mux
}

type PubSubMessage struct {
	Message      message `json:"message"`
	Subscription string  `json:"subscription"`
}

type message struct {
	Data       []byte     `json:"data,omitempty"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Type string `json:"msgType"`
}

func (s *Hl7Service) handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		respondJSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "request body is empty"},
		)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			respondJSON(
				w,
				http.StatusBadRequest,
				map[string]string{"error": p.Format("http.Request.Body.Close: %v", err)},
			)
			return
		}
	}()

	d := json.NewDecoder(r.Body)
	psMsg := PubSubMessage{}
	if err := d.Decode(&psMsg); err != nil {
		respondJSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": p.Format("json.Decoder.Decode: %v", err)},
		)
		return
	}

	hl7Path := string(psMsg.Message.Data)
	resp, err := s.msgService.Get(hl7Path).View("RAW_ONLY").Do()
	if err != nil {
		respondJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": p.Format("messages.Get: %v", err)},
		)
		return
	}
	hl7, err := base64.StdEncoding.DecodeString(resp.Data)
	if err != nil {
		respondJSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": p.Format("base64.DecodeString: %v", err)},
		)
		return
	}

	dataset := s.client.Dataset("methodist")

	msgType := psMsg.Message.Attributes.Type
	switch msgType {
	default:
		respondJSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": p.Format("unsupported message type: %s", msgType)},
		)
		return
	case "ADT":
		adt := ADT{}
		if err = json.Unmarshal(hl7, &adt); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("couldn't parse %s: %v", msgType, err)},
			)
			return
		}
		adt.MsgPath = hl7Path
		ins := dataset.Table("raw_adt").Inserter()
		if err = ins.Put(context.Background(), adt); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("client.Inserter.Put(%s): %v", msgType, err)},
			)
			return
		}
	case "ORM":
		orm := ORM{}
		if err = json.Unmarshal(hl7, &orm); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("couldn't parse %s: %v", msgType, err)},
			)
			return
		}
		orm.MsgPath = hl7Path
		ins := dataset.Table("raw_orm").Inserter()
		if err = ins.Put(context.Background(), orm); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("client.Inserter.Put(%s): %v", msgType, err)},
			)
			return
		}
	case "ORU":
		oru := ORU{}
		if err = json.Unmarshal(hl7, &oru); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("couldn't parse %s: %v", msgType, err)},
			)
			return
		}
		oru.MsgPath = hl7Path
		ins := dataset.Table("raw_oru").Inserter()
		if err = ins.Put(context.Background(), oru); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("client.Inserter.Put(%s): %v", msgType, err)},
			)
			return
		}
	case "MDM":
		mdm := MDM{}
		if err = json.Unmarshal(hl7, &mdm); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("couldn't parse %s: %v", msgType, err)},
			)
			return
		}
		mdm.MsgPath = hl7Path
		ins := dataset.Table("raw_mdm").Inserter()
		if err = ins.Put(context.Background(), mdm); err != nil {
			respondJSON(
				w,
				http.StatusInternalServerError,
				map[string]string{"error": p.Format("client.Inserter.Put(%s): %v", msgType, err)},
			)
			return
		}
	}

	respondJSON(w, http.StatusCreated, map[string]string{"success": p.Format("inserted a new %s", msgType)})
}
