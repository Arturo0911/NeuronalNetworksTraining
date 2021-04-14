import random

"""
Implementation of algorithm about the random asign 
in the weights to get values
"""


def equation():

    w1 = random.randint(1,9)
    w2 = random.randint(1,9)
    
    return w1, w2

def validation(w1,w2,x,y):

    return w1*x + w2*y == 10


if __name__ == '__main__':

    while True:
        
        w1, w2 = equation()
        
        if validation(w1,w2, 1.2,1.0):
            print(w1,w2)
            break



