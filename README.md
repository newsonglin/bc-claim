## Claim of BlockChain

A Node.js app to demonstrate claim related with hospital and insurer, this is based on fabric balance transfer example. I don't like write document as manay programmers. 
It's cruel to let a programmer write detail document. Howerver, I will still try to make detail enough to let you know how to start and play with it. So don't worry.


And I assume that you are farmiliar with hyperledger fabric (at least you know what it is),  and if you have already executed some of examples successfuly, that would be much easier for you to
execute this claim app.  

I finished this successfully on ubuntu 16.04,  if you use other OS, it should work, too.

### Prerequisites and setup:

* [Docker](https://www.docker.com/products/overview) - v1.12 or higher
* [Docker Compose](https://docs.docker.com/compose/overview/) - v1.8 or higher
* [Git client](https://git-scm.com/downloads) - needed for clone commands
* **Node.js** v8.4.0 or higher
* [Download Docker images](http://hyperledger-fabric.readthedocs.io/en/latest/samples.html#binaries)


```
cd bc-claim/claim/
```

Once you have completed the above setup, you will have provisioned a local network with the following docker container configuration:

* 2 CAs
* A SOLO orderer
* 4 peers (2 peers per Org)


## Running the claim program

There are two options available for running the claim 


### Option 1:

##### Terminal Window 1

* Launch the network using docker-compose

```
docker-compose -f artifacts/docker-compose.yaml up
```
##### Terminal Window 2

* Install the fabric-client and fabric-ca-client node modules

```
npm install
```

* Start the node app on PORT 4000

```
PORT=4000 node app
```

##### Terminal Window 3
* Install and Instantiate predefined chaincodes
```
cd bc-claim/claim

./init.sh

```


### Option 2:

##### Terminal Window 1

```
cd bc-claim/claim

./runApp.sh

```

* This lauches the required network on your local machine
* Installs the fabric-client and fabric-ca-client node modules
* And, starts the node app on PORT 4000


##### Terminal Window 2
* Install and Instantiate predefined chaincodes
```
cd bc-claim/claim

./init.sh

```


## Access the app

* http://localhost:4000/



