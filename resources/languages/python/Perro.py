#Esta la clase que se vuelve un tipo de variable
class Perro:
    #Atributos(Variables)
    nombre = "Firulais"
    edad = 0
    #Métodos o acciones(Procedimientos y funciones)
    def __init__(self):#Método constructor
        print("Perro creado")
    def ladrar(self):
        print("Guau Guau")
    def atacar(self):
        print("Grrrr")
    def decirNombre(self):
        nombre = "Bobby"
        print(nombre)#Bobby
        print(self.nombre)#Firulais

#Objeto1
perro1=Perro()
perro1.nombre="Bobby"
perro1.edad=3
perro1.ladrar()
perro1.atacar()
#Objecto2
perro2=Perro()
perro2.nombre="Hummer"
perro2.edad=1
perro2.ladrar()
perro2.atacar()
perro2.decirNombre()