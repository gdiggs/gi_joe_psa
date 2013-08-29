package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type PSAResponse struct {
	Text      string `json:"text"`
	Situation string `json:"situation"`
	Code_name string `json:"code_name"`
	Season    int    `json:"season"`
}

func (r PSAResponse) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		s = ""
		return
	}
	s = string(b)
	return
}

const (
	PCT_BATTLE_COMPLETE = 50
)

var (
	responses = []PSAResponse{
		{Code_name: "Alpine",
			Situation: "What to do when you're lost",
			Text:      "Stay calm. Think. Where did you see him last? Go back here. If he doesn't come back, ask a policeman for help.",
			Season:    1},

		{Code_name: "Barbeque",
			Situation: "What to do if your house is on fire",
			Text:      "Remember, if a fire breaks out in your home, always test the door first. If it's hot, find another exit or yell for help.",
			Season:    1},

		{Code_name: "Deep Six",
			Situation: "Don't swim when it's storming",
			Text:      "At even the hint of a thunderstorm, get right out of the water.",
			Season:    1},

		{Code_name: "Mutt",
			Situation: "Don't pet strange dogs",
			Text:      "Don't run! Walk away slowly. Never try to pet an animal you don't know. He might be lost, sick or scared. If we don't know, we leave 'em alone.",
			Season:    1},

		{Code_name: "Quick Kick",
			Situation: "Don't be in a hurry to build your tree house",
			Text:      "Remember, anything worth doing is worth planning.",
			Season:    1},

		{Code_name: "Spirit",
			Situation: "What to do if you catch on fire",
			Text:      "Remember, if your clothes catch on fire, rap yourself in a rug or blanket (and roll on the ground to smother the flames).",
			Season:    1},

		{Code_name: "Torpedo",
			Situation: "How to tread water",
			Text:      "Open and close your legs like a scissor. Keep up a steady rhythm. Now cup your hands downward and move them in a figure eight motion... Never play around water alone.",
			Season:    1},

		{Code_name: "Cutter",
			Situation: "Have a ump ref your baseball games",
			Text:      "Well, fighting won't stop it. When people disagree sometimes they need someone who's not involved to settle things.",
			Season:    1},

		{Code_name: "Flint",
			Situation: "Don't get mad at your goalie for letting in a bad goal",
			Text:      "Will yelling at Billy help?... Look, if you want to play your best you got to play like a team. Remember, you need teamwork to win, not arguments.",
			Season:    1},

		{Code_name: "Dusty",
			Situation: "Put reflectors on your bike or be runover",
			Text:      "And I couldn't see you! No wonder, you don't have reflectors. They tell drivers where you are... Remember, if you have to ride when it's getting dark, have the right equipment.",
			Season:    1},

		{Code_name: "Scarlett",
			Situation: "You can learn to water ski if you keep trying",
			Text:      "That's because you quit trying... You'll never win if you give in.",
			Season:    1},

		{Code_name: "Roadblock ",
			Situation: "Don't tell a stranger where you live",
			Text:      "Remember, never tell anyone you're home alone and never give anyone your address.",
			Season:    1},

		{Code_name: "Flint",
			Situation: "Don't blame your brother",
			Text:      "Anyone can help an accident, but lying makes it worse... Face up to what you've done. Don't take the easy way out. Remember, it's better to tell the truth.",
			Season:    1},

		{Code_name: "Barbeque",
			Situation: "Don't pull the fire alarm unless there's a fire",
			Text:      "Remember, a firefighter's job is to fight fires - not answer false alarms.",
			Season:    1},

		{Code_name: "Footloose",
			Situation: "How to stop a nose bleed",
			Text:      "Pinch your nose closed and lean forward. If it doesn't stop in 5 minutes, pack your nose with gauze and pinch it closed for ten more minutes. If its still bleeding, then see a doctor",
			Season:    1},

		{Code_name: "Snow Job",
			Situation: "Don't skate on thin ice",
			Text:      "Remember, frozen ponds and rivers may not be totally frozen",
			Season:    1},

		{Code_name: "Doc",
			Situation: "Don't take drugs without your parents there.",
			Text:      "Never take medicine without a grown up present. You could do more harm then good... If you can, wait for your parents. Of if it's serious, ask a neighbor for help.",
			Season:    1},

		{Code_name: "Lady Jaye",
			Situation: "It's okay to be a chicken if your smart",
			Text:      "There's nothing chicken about being smart. If you stop and think there's almost always a better way. ",
			Season:    1},

		{Code_name: "Ripcord",
			Situation: "Maybe you stink at baseball because you need glasses",
			Text:      "Having your eyes tested may clear things up. Don't avoid a problem. Meet it ( and beat it).",
			Season:    1},

		{Code_name: "Shipwreck",
			Situation: "Don't run away from home",
			Text:      "Isn't it better to try to solve problems instead of runnin' away from 'em?... And remember, running away (leads nowhere).",
			Season:    1},

		{Code_name: "Recondo",
			Situation: "Don't hide in the frige",
			Text:      "Remember, never get in anything that could close up and trap you. ",
			Season:    1},

		{Code_name: "Airtight",
			Situation: "What to do if someone passes out",
			Text:      "Never lift the head of someone who's fainted... Keep him flat and brace his legs. Now, loosen his clothes and use a wet cloth. ",
			Season:    1},

		{Code_name: "Roadblock",
			Situation: "Don't jump your bike over downed power lines ",
			Text:      "Remember, don't play around electric wires or you could be playing with fire. ",
			Season:    1},

		{Code_name: "Blowtorch",
			Situation: "Don't call the fire department from a burning building ",
			Text:      "If there's a real fire in your house, your first job is to escape immediately. Fire spreads quickly. Call the fire department from outside the house.",
			Season:    1},

		{Code_name: "Gung-Ho",
			Situation: "Girls can skateboard better too ",
			Text:      "Don't judge people 'til you give them a chance.",
			Season:    1},

		{Code_name: "Shipwreck",
			Situation: "Stealing is wrong ",
			Text:      "Remember, taking something that isn't yours just isn't right.",
			Season:    1},

		{Code_name: "Will Bill",
			Situation: "Don't go with strangers ",
			Text:      "Well, just don't do what a stranger says. Check it out with an adult you know. Remember, a stranger (can mean danger).",
			Season:    1},

		{Code_name: "Deep Six",
			Situation: "Wear a life jacket",
			Text:      "A life jacket is good protection (like seat belts in a car).",
			Season:    1},

		{Code_name: "Leatherneck",
			Situation: "Don't stay out all day in the sun or you'll turn into a lobster.",
			Text:      "A bad sunburn can make you sick or even put you in the hospital.",
			Season:    2},

		{Code_name: "Lifeline",
			Situation: "Don't eat junk food",
			Text:      "A candy bar might give you a quick boost of energy, but after 20 minutes, you'll feel run down. So let's eat smart!",
			Season:    2},

		{Code_name: "Spirit",
			Situation: "Blind kids can find lost kittens too",
			Text:      "Having a handicap doesn't mean you're helpless.",
			Season:    2},

		{Code_name: "Cross Country & Beach Head",
			Situation: "Protect your head",
			Text:      "These ATVs have as much power as a real motorcycle. You gotta respect 'em. (And wear a helmet for protection).",
			Season:    2},

		{Code_name: "Dial Tone",
			Situation: "Don't paint your bike in the garage",
			Text:      "All paints, and especially spray paints, have poisonous gases in 'em. If you breath too much you can get very sick... Always read the label carefully and check for warnings before you start any job.",
			Season:    2},

		{Code_name: "General Hawk",
			Situation: "Don't try to beat the train ",
			Text:      "That's a dangerous game you two were playing. Listen to your friend Chris. Those gates are provided as a warning to let you known that it's not safe to cross. ",
			Season:    2}}

	max_responses = len(responses)
)

func v1handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, responses[rand.Intn(max_responses)])
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", v1handler)
	http.HandleFunc("/v1", v1handler)
	http.HandleFunc("/v1.json", v1handler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
