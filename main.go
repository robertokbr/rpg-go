package main

import (
	"fmt"
	"time"
)

func reader() string {
	var data string
	fmt.Scan(&data)
	return data
}

func printer(message string) {
	fmt.Println(message)
}

// declaracao de uma classe
type Player struct {
	Name   string
	Health int
	Damage int
}

func NewPlayer(name string) Player {
	return Player{
		Name:   name,
		Health: 5,
		Damage: 1,
	}
}

type NPC struct {
	Name   string
	Health int
	Damage int
}

func NewNPC(name string, damage, health int) NPC {
	return NPC{
		Name:   name,
		Damage: damage,
		Health: health,
	}
}

func await() {
	time.Sleep(time.Second * 2)
}

func main() {
	var player Player

	printer("Ola jogador, qual é seu nome?")
	playerName := reader()
	printer("Prazer em lhe conhecer " + playerName)

	printer("Você gostaria de jogar um joguin?")
	printer("[s/n]")
	playerResponse := reader()

	player = NewPlayer(playerName)

	if playerResponse == "n" {
		printer("Então vai se lascar seu lixo humano!")
		return
	}

	printer("Pois bem, vamos começar então!")

	await()

	printer("[Você estava em uma rua escura, a tempos você não se sentia vivo...]")

	await()

	printer("[E no meio da escuridão, você se deparou com um cachorro falante...]")

	dog := `
			/ \__
			(    @\___
			/         O
			/   (_____/
			/_____/   U
	`

	printer(dog)

	printer("[dog] Ola humano, você esta feliz em me ver?")

	printer("[s/n]")

	playerResponse = reader()

	if playerResponse == "s" {
		printer("[dog] Pois bem, eu também estou feliz em lhe ver!")
	} else {
		printer("[dog] É realmente uma pena, pois eu estou feliz em lhe ver!")

		await()

		player.Health -= 1

		printer("[Você magoou os sentimentos do dog magico...]")
		printer("[O universo lhe atingiu com uma rajada de vento e você perdeu 1 de vida]")
		message := fmt.Sprintf("[Você agora possui apenas %v de vida]", player.Health)
		printer(message)
	}

	await()

	printer("[silencio...]")

	await()

	printer("[dog] Prepare-se, vem vindo encrenca da braba")

	badCat := NewNPC("bad cat", 1, 3)

	cat := `
	../\_/\  
	.( o.o ) 
	..> ^ <
	`

	printer(cat)

	badCatMessage := fmt.Sprintf("[%v] ora ora ora, veja quem eu encontrei...", badCat.Name)
	await()

	printer(badCatMessage)

	badCatMessage = fmt.Sprintf("[%v] preparem-se para morrer!...", badCat.Name)

	await()

	printer(badCatMessage)

	message := fmt.Sprintf("[%v lhe atacou e lhe inflingiu %v de dano]", badCat.Name, badCat.Damage)

	printer(message)

	player.Health -= badCat.Damage

	printer("O que você pretende fazer?")
	printer("[correr/revidar]")

	choice := reader()

	printer("Voce escolheu " + choice)
}
