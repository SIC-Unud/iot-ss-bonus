#include <ESP8266WiFi.h>

#include <DNSServer.h>
#include <ESP8266WebServer.h>
#include <WiFiManager.h>

#include <WiFiClient.h>
#include <ESP8266HTTPClient.h>


#include "DHT.h"
#define DHTPIN 5
#define DHTTYPE DHT11

WiFiManager wifiManager;
DHT dht(DHTPIN, DHTTYPE);

const char *serverName = "server_endpoint";
unsigned long lastTime = 0;
unsigned long timerDelay = 300000;

void setup() {
  Serial.begin(115200);

  // comment 3 lines below to make the ESP auto connect with connected WiFi before
  // and remove the timeout
  WiFi.mode(WIFI_STA);
  wifiManager.resetSettings();
  wifiManager.setTimeout(180);

  if(!wifiManager.autoConnect("IoT Sharing Session", "iot12345")) {
    Serial.println("Failed to connect and hit timeout");
    delay(3000);
    ESP.reset();
    delay(5000);
  }

  Serial.println("Connected!");
  dht.begin();

}

void loop() {
  if((millis() - lastTime) > timerDelay) {
    if(WiFi.status() == WL_CONNECTED) {
      WiFiClient client;
      HTTPClient http;

      http.begin(client, serverName);
      
      float h = dht.readHumidity();
      float t = dht.readTemperature();

      if(isnan(h) || isnan(t)) {
        Serial.println("Failed to read fron DHT sensor!");
      }

      Serial.print("Temperature: ");
      Serial.print(t);
      Serial.print("°C  Humidity: ");
      Serial.print(h);
      Serial.println("%");

      http.addHeader("Content-Type", "application/json");
      String json = "{\"temp\": " + (String)t + ", \"humid\": " + (String)h + "}";
      int httpResponseCode = http.POST(json);

      Serial.print("HTTP Response code: ");
      Serial.println(httpResponseCode);

      http.end();
    } else {
      Serial.println("WiFi Disconnected");
    }
    lastTime = millis();
  }

}
