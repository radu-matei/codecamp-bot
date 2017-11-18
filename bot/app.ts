import * as restify from 'restify';
import * as builder from 'botbuilder';
import * as os from 'os';
import * as http from 'http'

const connector = new builder.ChatConnector({
    appId: process.env.MICROSOFT_APP_ID,
    appPassword: process.env.MICROSOFT_APP_PASSWORD
});

const server = restify.createServer();
server.post('/api/messages', connector.listen());
server.get('/healthz', (request, response) => {
    response.send(200);
})

server.listen(3978, () => console.log(`${server.name} listening to ${server.url}`));

var bot = new builder.UniversalBot(connector, (session: builder.Session) => {
    session.send(`I am ${os.hostname}. You said: ${session.message.text}`)
});

var recognizer = new builder.LuisRecognizer(process.env.LUIS_URI);
bot.recognizer(recognizer);


bot.dialog('GetInformation', function (session) {
    var options = {
        host: 'go-client',
        port: 80,
        path: '/get/cluster'
    };

    var body = "Hi, Sofia!";
    http.get(options, response => {
        response.on('data', data => {
            body += data
        });
        response.on('end', ()=> {
            session.say(body, body)     
        })
    })
}).triggerAction({
    matches: 'GetInformation'
});

bot.dialog('Deploy', function (session) {
    var options = {
        host: 'go-client',
        port: 80,
        path: '/create'
    };

    var body = "";
    http.get(options, response => {
        response.on('data', data => {
            body += data
        });
        response.on('end', ()=> {
            session.say(body, body)     
        })
    })

}).triggerAction({
    matches: 'Deploy'
});


bot.dialog('Update', function (session) {
    var options = {
        host: 'go-client',
        port: 80,
        path: '/update'
    };

    var body = "I am updating the application. ";
    http.get(options, response => {
        response.on('data', data => {
            body += data
        });
        response.on('end', ()=> {
            session.say(body, body)     
        })
    })

}).triggerAction({
    matches: 'Update'
});