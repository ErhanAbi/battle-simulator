### Some Sort of Hero

#### The story

Once upon a time there was a great hero, named however you feel like naming him, with some strengths and weaknesses, as all heroes have.
After battling all kinds of monsters for more than a hundred years, the hero now has the following stats that are generated when he is summoned(or initialized if that sounds better to you):

- Health: 70 - 100
- Strength: 70 - 80
- Defence: 45 – 55
- Speed: 40 – 50
- Luck: 10% - 30% (0% means no luck, 100% lucky all the time)

He also possesses 2 skills:

- Critical Strike: Strike twice while it’s his turn to attack; there’s a 10% chance he’ll use this skill every time he attacks. If he uses this skill then there is also a 1% chance that he’ll strike three times instead of two.
- Resilience: Takes only half of the usual damage when an enemy attacks; there’s a 20% chance he’ll use this skill when he defends but this skill cannot be used two turns in a row.

#### Gameplay

As the hero walks the whimsical forests of The Terminal Valley, he encounters nefarious villains, with the following properties:

- Health: 60 - 90
- Strength: 60 - 90
- Defence: 40 – 60
- Speed: 40 – 60
- Luck: 25% - 40%

You’ll have to simulate a battle between the hero and a nefarious villain, either at command line or using a web browser. On every battle, the hero and the villain must be initialized with random properties, within their ranges. The first attack is done by the player with the higher speed. If both players have the same speed, then the attack is carried on by the player with the highest luck. After an attack, the players switch roles: the attacker now defends and the defender now attacks.

The damage done by the attacker is calculated with the following formula:
Damage = Attacker strength – Defender defence
The damage is subtracted from the defender’s health. An attacker can miss their hit and do no damage if the defender gets lucky that turn.
The hero’s skills occur randomly, based on their chances, so take them into account on each turn.

#### Game over

The game ends when one of the players remain without health or the number of turns reaches 20. The application must output the results each turn: what happened, which skills were used (if any), the damage done, defender’s health left. If we have a winner before the maximum number of rounds is reached, he must be declared.

#### Rules

- Write code in plain Go/Python/NodeJS (you are free to use 3rd parties like UI libs / frameworks / etc)
- Make sure your application is decoupled, code reusable and scalable. For example, can a new skill easily be added to our hero?
- Is your code bug-free and tested?
- There’s no time limit, take your time for the best approach you can think of.
  Sharing the solution The code for the implemented solution should be managed by a source control system, and the link to a public repository shared over e-mail.
