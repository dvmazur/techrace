import random

# Choice names
r = "Rock"
p = "Paper"
sc = "Scissors"
sp = "Spock"
l = "Lizard"

# Who beats whom
beat_table = {
    r : [ l, sc ],
    p : [ sp, r ],
    sc : [ p, l ],
    sp : [ r, sc ],
    l : [ p, sp ] 
}

# How do they beat each other?
beat_name = {
    sc+p : "Scissors cut Paper!",
    p+r : "Paper covers Rock!",
    r+l : "Rock crushes Lizard!",
    l+sp : "Lizard poisons Spock!",
    sp+sc: "Spock smashes Scissors!",
    sc+l : "Scissors decapitates Lizard!",
    l+p : "Lizard etas Paper!",
    p+sp : "Paper disproves Spock!",
    sp+r : "Spock vaporizes Rock!",
    r+sc : "Rock crushes Scissors!"
}

full_names = [ r, p, sc, sp, l ]
short_names = [ "R", "P", "Sc", "Sp", "L" ]

short_to_full = {
    "R" : r,
    "P" : p,
    "Sc": sc,
    "Sp": sp,
    "L": l
}

while True:
    print("Choose your hero!")
    print("(R)ock, (P)aper, (Sc)issors, (Sp)ock, (L)izard")

    player_choice = input()

    if player_choice not in full_names:
        if player_choice in short_names:
            player_choice = short_to_full[player_choice]
        else:
            print("Oh no! Seems like you misspelled something. Try again!")
            continue

    computer_choice = random.choice(full_names)

    print("Computer randomly chose {}!".format(computer_choice))

    if computer_choice in beat_table[player_choice]:
        print(beat_name[player_choice+computer_choice])
        print("You won! Let's play more? (Y)es/(N)o")
    elif player_choice in beat_table[computer_choice]:
        print(beat_name[computer_choice+player_choice])
        print("You lost! Wanna take a revenge? (Y)es/(N)o")
    else:
        print("Seems like it's a draw. Want to try once again? (Y)es/(N)o")

    go_on = input()

    if(go_on == "Y" or go_on == "Yes"):
        print("That's noice! Let's play!")
    elif(go_on == "N" or go_on == "No"):
        print("That's a shame you don't want to play... Cause I do!")
        exit()
    else:
        print("I didn't quite get what you've said, but I guess you want to play more!")