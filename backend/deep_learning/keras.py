# import sys

# def main():
#     if len(sys.argv) > 1:
#         print(f"This is Keras Hello, {sys.argv[1]}!")
#     else:
#         print("Mr.Keras, Hello, World!")

# if __name__ == "__main__":
#     main()

import numpy as np
import sys

def main():
    if len(sys.argv) > 1:
        value = float(sys.argv[1])
    else:
        value = 10.0

    arr = np.array([value] * 5)
    result = np.mean(arr)
    
    print(f"Array: {arr}")
    print(f"Mean: {result}")

if __name__ == "__main__":
    main()

