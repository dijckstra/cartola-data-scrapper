package model

// Scout contains performance indicators for a player.
type Scout struct {
	A  int32 // Assistência
	CA int32 // Cartão Amarelo
	CV int32 // Cartão Vermelho
	DD int32 // Defesa Difícil
	DP int32 // Defesa de Pênalti
	FC int32 // Falta Cometida
	FD int32 // Finalização Defendida
	FF int32 // Finalização Fora
	FS int32 // Falta Sofrida
	FT int32 // Finalização na Trave
	G  int32 // Gol
	GC int32 // Gol Contra
	GS int32 // Gol Sofrido
	I  int32 // Impedimento
	PE int32 // Passe Errado
	PP int32 // Pênalti Perdido
	RB int32 // Roubada de Bola
	SG int32 // Jogos sem Sofrer Gols
}
