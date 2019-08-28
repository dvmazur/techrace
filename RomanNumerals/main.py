def numeral_to_roman(number):
    if not 0 < number < 4000:
        raise RuntimeError
    nums = (1000, 900,  500, 400, 100,  90, 50,  40, 10,  9,   5,  4,   1)
    roman = ('M',  'CM', 'D', 'CD','C', 'XC','L','XL','X','IX','V','IV','I')
    result = []
    for i in range(len(nums)):
        count = int(number / nums[i])
        result.append(roman[i] * count)
        number -= nums[i] * count
    return ''.join(result)

while True:
    print("Enter your numeral!")
    try:
        numeral = int(input())
        roman = numeral_to_roman(numeral)
    except ValueError:
        print("You should enter an integer!")
        continue
    except RuntimeError:
        print("Number must be between 1 and 3999!")
        continue
    else:
        print("Roman numer for {} is {}".format(numeral, roman))