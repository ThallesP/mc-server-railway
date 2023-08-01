<h1 align="center">Minecraft Server for Java</h1>

Railway is great for hosting your applications, but what if we tried to host a Minecraft Server? The benefits are incredible:

- You only pay for what you use, so it's usually cheap
- No more upgrading your server resources
- Super stable
- Railway Support is the best

# Enough talk, how do I deploy this?
It's simple, follow me:

1. Create an Ngrok Account on ngrok.com and copy your Authtoken.
2. Click on "Deploy Now" button in this Railway's template
3. Fill the necessary info (Accept EULA and enter your Authtoken)

After this, your Minecraft Server will be deployed.

# How do I connect to it?

Click on the "Ngrok" service, go to logs and the IP will be there.

# How I can configure it?
As we use the [itzg/minecraft-server](https://hub.docker.com/r/itzg/minecraft-server) docker image, you can practically configure anything you want through environment variables.

See their [docs](https://docker-minecraft-server.readthedocs.io/en/latest/variables/) for more info.
