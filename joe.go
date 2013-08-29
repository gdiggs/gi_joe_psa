package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PSAResponse struct {
	text      string
	situation string
	code_name string
	season    int
}

const (
	PCT_BATTLE_COMPLETE = 50
)

var (
	responses = []PSAResponse{
		{code_name: "Alpine",
			situation: "What to do when you're lost",
			text:      "Stay calm. Think. Where did you see him last? Go back here. If he doesn't come back, ask a policeman for help.",
			season:    1},

		{code_name: "Barbeque",
			situation: "What to do if your house is on fire",
			text:      "Remember, if a fire breaks out in your home, always test the door first. If it's hot, find another exit or yell for help.",
			season:    1},

		{code_name: "Deep Six",
			situation: "Don't swim when it's storming",
			text:      "At even the hint of a thunderstorm, get right out of the water.",
			season:    1},

		{code_name: "Mutt",
			situation: "Don't pet strange dogs",
			text:      "Don't run! Walk away slowly. Never try to pet an animal you don't know. He might be lost, sick or scared. If we don't know, we leave 'em alone.",
			season:    1},

		{code_name: "Quick Kick",
			situation: "Don't be in a hurry to build your tree house",
			text:      "Remember, anything worth doing is worth planning.",
			season:    1},

		{code_name: "Spirit",
			situation: "What to do if you catch on fire",
			text:      "Remember, if your clothes catch on fire, rap yourself in a rug or blanket (and roll on the ground to smother the flames).",
			season:    1},

		{code_name: "Torpedo",
			situation: "How to tread water",
			text:      "Open and close your legs like a scissor. Keep up a steady rhythm. Now cup your hands downward and move them in a figure eight motion... Never play around water alone.",
			season:    1},

		{code_name: "Cutter",
			situation: "Have a ump ref your baseball games",
			text:      "Well, fighting won't stop it. When people disagree sometimes they need someone who's not involved to settle things.",
			season:    1},

		{code_name: "Flint",
			situation: "Don't get mad at your goalie for letting in a bad goal",
			text:      "Will yelling at Billy help?... Look, if you want to play your best you got to play like a team. Remember, you need teamwork to win, not arguments.",
			season:    1},

		{code_name: "Dusty",
			situation: "Put reflectors on your bike or be runover",
			text:      "And I couldn't see you! No wonder, you don't have reflectors. They tell drivers where you are... Remember, if you have to ride when it's getting dark, have the right equipment.",
			season:    1},

		{code_name: "Scarlett",
			situation: "You can learn to water ski if you keep trying",
			text:      "That's because you quit trying... You'll never win if you give in.",
			season:    1},

		{code_name: "Roadblock ",
			situation: "Don't tell a stranger where you live",
			text:      "Remember, never tell anyone you're home alone and never give anyone your address.",
			season:    1},

		{code_name: "Flint",
			situation: "Don't blame your brother",
			text:      "Anyone can help an accident, but lying makes it worse... Face up to what you've done. Don't take the easy way out. Remember, it's better to tell the truth.",
			season:    1},

		{code_name: "Barbeque",
			situation: "Don't pull the fire alarm unless there's a fire",
			text:      "Remember, a firefighter's job is to fight fires - not answer false alarms.",
			season:    1},

		{code_name: "Footloose",
			situation: "How to stop a nose bleed",
			text:      "Pinch your nose closed and lean forward. If it doesn't stop in 5 minutes, pack your nose with gauze and pinch it closed for ten more minutes. If its still bleeding, then see a doctor",
			season:    1},

		{code_name: "Snow Job",
			situation: "Don't skate on thin ice",
			text:      "Remember, frozen ponds and rivers may not be totally frozen",
			season:    1},

		{code_name: "Doc",
			situation: "Don't take drugs without your parents there.",
			text:      "Never take medicine without a grown up present. You could do more harm then good... If you can, wait for your parents. Of if it's serious, ask a neighbor for help.",
			season:    1},

		{code_name: "Lady Jaye",
			situation: "It's okay to be a chicken if your smart",
			text:      "There's nothing chicken about being smart. If you stop and think there's almost always a better way. ",
			season:    1},

		{code_name: "Ripcord",
			situation: "Maybe you stink at baseball because you need glasses",
			text:      "Having your eyes tested may clear things up. Don't avoid a problem. Meet it ( and beat it).",
			season:    1},

		{code_name: "Shipwreck",
			situation: "Don't run away from home",
			text:      "Isn't it better to try to solve problems instead of runnin' away from 'em?... And remember, running away (leads nowhere).",
			season:    1},

		{code_name: "Recondo",
			situation: "Don't hide in the frige",
			text:      "Remember, never get in anything that could close up and trap you. ",
			season:    1},

		{code_name: "Airtight",
			situation: "What to do if someone passes out",
			text:      "Never lift the head of someone who's fainted... Keep him flat and brace his legs. Now, loosen his clothes and use a wet cloth. ",
			season:    1},

		{code_name: "Roadblock",
			situation: "Don't jump your bike over downed power lines ",
			text:      "Remember, don't play around electric wires or you could be playing with fire. ",
			season:    1},

		{code_name: "Blowtorch",
			situation: "Don't call the fire department from a burning building ",
			text:      "If there's a real fire in your house, your first job is to escape immediately. Fire spreads quickly. Call the fire department from outside the house.",
			season:    1},

		{code_name: "Gung-Ho",
			situation: "Girls can skateboard better too ",
			text:      "Don't judge people 'til you give them a chance.",
			season:    1},

		{code_name: "Shipwreck",
			situation: "Stealing is wrong ",
			text:      "Remember, taking something that isn't yours just isn't right.",
			season:    1},

		{code_name: "Will Bill",
			situation: "Don't go with strangers ",
			text:      "Well, just don't do what a stranger says. Check it out with an adult you know. Remember, a stranger (can mean danger).",
			season:    1},

		{code_name: "Deep Six",
			situation: "Wear a life jacket",
			text:      "A life jacket is good protection (like seat belts in a car).",
			season:    1},

		{code_name: "Leatherneck",
			situation: "Don't stay out all day in the sun or you'll turn into a lobster.",
			text:      "A bad sunburn can make you sick or even put you in the hospital.",
			season:    2},

		{code_name: "Lifeline",
			situation: "Don't eat junk food",
			text:      "A candy bar might give you a quick boost of energy, but after 20 minutes, you'll feel run down. So let's eat smart!",
			season:    2},

		{code_name: "Spirit",
			situation: "Blind kids can find lost kittens too",
			text:      "Having a handicap doesn't mean you're helpless.",
			season:    2},

		{code_name: "Cross Country & Beach Head",
			situation: "Protect your head",
			text:      "These ATVs have as much power as a real motorcycle. You gotta respect 'em. (And wear a helmet for protection).",
			season:    2},

		{code_name: "Dial Tone",
			situation: "Don't paint your bike in the garage",
			text:      "All paints, and especially spray paints, have poisonous gases in 'em. If you breath too much you can get very sick... Always read the label carefully and check for warnings before you start any job.",
			season:    2},

		{code_name: "General Hawk",
			situation: "Don't try to beat the train ",
			text:      "That's a dangerous game you two were playing. Listen to your friend Chris. Those gates are provided as a warning to let you known that it's not safe to cross. ",
			season:    2}}

	max_responses = len(responses)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(responses[rand.Intn(max_responses)])
}
