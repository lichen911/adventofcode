
class MarbleGame(object):
    def __init__(self, num_players, last_marble):
        self.num_players = num_players
        self.last_marble = last_marble
        self.circle_list = [0]
        self.player_turn = 0
        self.player_list = [0 for _ in range(num_players)]
        self.current_marble_idx = 0
        self.highest_played_marble = 0

    def play_marble(self, new_marble):
        if self.player_turn == self.num_players:
            self.player_turn = 1
        else:
            self.player_turn += 1

        # test if this player scores
        if new_marble % 23 == 0:
            self.add_player_score(self.player_turn, new_marble)
            removed_marble_idx = (self.current_marble_idx - 7) % len(self.circle_list)
            removed_marble_value = self.circle_list.pop(removed_marble_idx)
            self.add_player_score(self.player_turn, removed_marble_value)
            self.current_marble_idx = removed_marble_idx
        else:
            # if not then do a normal turn
            if len(self.circle_list) == 1:
                # print('first turn')
                self.circle_list = [0, 1]
                self.current_marble_idx = 1
            elif self.current_marble_idx + 2 == len(self.circle_list):
                # print('add to end')
                self.current_marble_idx = len(self.circle_list)
                self.circle_list.append(new_marble)
            else:
                # print('wrap around')
                new_current_marble = (self.current_marble_idx + 2) % len(self.circle_list)
                self.circle_list.insert(new_current_marble, new_marble)
                self.current_marble_idx = new_current_marble

    def get_next_marble(self):
        self.highest_played_marble += 1

        if self.highest_played_marble > self.last_marble:
            return None
        else:
            return self.highest_played_marble

    def get_player_score(self, player):
        return self.player_list[player-1]

    def add_player_score(self, player, value):
        self.player_list[player - 1] += value

    def get_highest_score(self):
        return max(self.player_list)

    def __str__(self):
        new_list = [str(i) for i in self.circle_list]
        new_list[self.current_marble_idx] = '({})'.format(new_list[self.current_marble_idx])

        if self.player_turn == 0:
            new_list.insert(0, '[-]')
        else:
            new_list.insert(0, '[{}]'.format(self.player_turn))

        # print('current_marble: {}, highest_played_marble: {}, circle_list: {}'.format(self.current_marble, self.highest_played_marble, self.circle_list))
        return ' '.join(new_list)


def play_game(num_players, last_marble):
    game_board = MarbleGame(num_players, last_marble)
    # print(game_board)

    while True:
        next_marble = game_board.get_next_marble()
        if next_marble:
            game_board.play_marble(next_marble)
            # print(max(game_board.player_list))
            # print(game_board)
        else:
            break

    print('Highest Elf score: {}'.format(game_board.get_highest_score()))


def main():
    # games_to_play = [(9, 25)]
    # games_to_play = [(10, 1000)]
    games_to_play = [(9, 25), (10, 1618), (13, 7999), (17, 1104), (21, 6111), (30, 5807), (431, 70950), (431, 7095000)]
    # games_to_play = [(10, 10), (10, 1000), (10, 100000)]

    for game in games_to_play:
        print('\nplaying game: {}'.format(game))
        play_game(*game)


if __name__ == "__main__":
    main()
