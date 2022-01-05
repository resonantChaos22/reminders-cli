const express = require("express");
const bodyParser = require("body-parser");
const notifier = require("node-notifier");

const app = express();
const port = process.env.PORT || 9000;

app.use(bodyParser.json());

// Endpoints
app.get("/health", (req, res) =>
  res.status(200).send("Notifier service is running")
);
app.post("/notify", async (req, res) => {
  await notify(req.body, (reply) => res.send(reply));
});

const notify = ({ title, message }, cb) => {
  notifier.notify(
    {
      title: title || "Unknown Title",
      message: message || "Unknown Message",
      sound: true,
      icon: "Terminal Icon",
      wait: true,
      reply: true,
      closeLabel: "Completed?",
      timeout: 15,
    },
    (err, response, reply) => {
      reply = {
        deliveredAt: new Date().toUTCString(),
        activationType: "closed",
        activationAt: new Date().toUTCString(),
        activationValue: "Completed?",
      };

      cb(reply);
    }
  );
};

app.listen(port, () =>
  console.log(`Server is up and running at port: ${port}`)
);
