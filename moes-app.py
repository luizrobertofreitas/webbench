# pip install japronto

from japronto import Application

def factorial(number):
    return 1 if number == 0 else number*factorial(number-1)

def fac (request):
    pathParam = request.match_dict
    n = int(pathParam.get('n'))
    soma =  factorial(n)
    return request.Response(text=str(soma))

app = Application()
app.router.add_route('/{n}',fac)
app.run(debug=False, port=8090)
