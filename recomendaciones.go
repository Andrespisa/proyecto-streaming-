package main

import "time"

type Recomendaciones struct {
	IDRecomendacion    string
	IDPerfil           string
	IDContenido        string
	FechaRecomendacion time.Time
}
