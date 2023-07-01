# Objective

<p>The client sends info(name and age) to the server, The server processes this information and informs the user whether he is eligible to apply for a drivers license or not</p>

## Server

<p> The server listens continously on the specified port for information from the client. Once the information is written into the buffer by the client it is read by the server. The information from the client comes in json format with a Name and Age field. The json file is unmarshelled first and the age is processed. According to the age input by the client appropriate response is given to the client and his/her session is terminated</p>

## Client

<p> The client creates a json with name and age , then adds the json into the buffer. The server processed this info and sends an appropriate response</p>