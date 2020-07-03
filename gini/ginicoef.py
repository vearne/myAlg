import numpy as np

def gini_coef(wealths):
    cum_wealths = np.cumsum(sorted(np.append(wealths, 0)))
    print("cum_wealths", cum_wealths)
    sum_wealths = cum_wealths[-1]
    xarray = np.array(range(0, len(cum_wealths))) / np.float(len(cum_wealths)-1)
    print("xarray", xarray)
    yarray = cum_wealths / sum_wealths
    print("yarray", yarray)
    B = np.trapz(yarray, x=xarray)
    A = 0.5 - B
    return A / (A+B)


if __name__ == "__main__":
    wealths = [10, 1000, 100, 20, 50]
    print(gini_coef(wealths))
