import adafruit_dht
import board

dhtDevice = adafruit_dht.DHT11(4)

print (board.D17)

temp = dhtDevice.temperature
humid = dhtDevice.humidity

print ("Temp", temp, "Humidity", humid)