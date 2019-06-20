# Business House Board Game

## Problem statement

Business House is a board game, popularly known as Monopoly as well. In this problem at hand we are going to simluate the same but with minimal features, all of which are described below:

Number of Players - 2 or more.
Each player is positioned at the starting box on the game board, with an initial balance of 2000 each. They get a turn to throw a couple of dice, so they may get any value between 2 and 12(inclusive). Based on the value of the dice face values, the player moves along the boxes on the board. Each box can be one of three types.

    - JAIL - The player has to pay 150 rupees(ideally to a bank but in this case we can consider them to just get debited from the player's account)
    - TREASURE -  The player is rewarded with 200 rupees.
    - HOTEL - If the hotel is not owned by another player and the player has enough balance ,then he has to buy the hotel for it's cost 150 rupees.
    - EMPTY - The player has to do nothing

At the end of 10 rounds of throwing the dice, the player with the highest balance is winner, so display the players with their balance in decresiing order of balance.


## Input 

Number of players -> for example 2.
Dice throws -> [2,12,3,5,10,6,11,4,2,5,3,9,8,4,6,10,4,5,12,10]
Board layout -> [E,E,H,H,E,T,H,E,H,H,E,E,E,J,T,E,H]


## Output

Player 1 - 1200
Player 2 - 800