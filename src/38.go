from dataclasses import dataclass
import math

@dataclass
class SoccerBall:
    x: float = 0.0
    y: float = 0.0
    vx: float = 0.0
    vy: float = 0.0

def find_closest_ball(soccer_balls):
    closest_x, closest_y, min_distance = math.inf, math.inf, 1e9
    for ball in soccer_balls:
        distance = ((ball.x - closest_x) ** 2 + (ball.y - closest_y) ** 2) ** 0.5
        if distance < min_distance:
            min_distance = distance
            closest_x, closest_y = ball.x, ball.y
    return SoccerBall(closest_x, closest_y, 0.0, 0.0)

if __name__ == "__main__":
    # Example soccer balls data
    example_balls = [
        SoccerBall(1.0, 2.5, 0.0, 0.0),
        SoccerBall(-3.0, -4.5, 0.0, 0.0),
        SoccerBall(6.0, 8.0, 0.0, 0.0)
    ]
    closest_ball = find_closest_ball(example_balls)
    print("Closest ball:", closest_ball.x, closest_ball.y)
