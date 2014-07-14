# -*- coding: utf-8 -*-
import requests
import ujson as json


class SlackBot(object):
    def __init__(self, token, team, channel=None):
        self.token = token
        self.channel = channel
        self.api_url = 'https://%s.slack.com/services/hooks/slackbot' % team

    def send(self, msg, channel=None):
        if not channel and self.channel:
            channel = self.channel
        elif not (channel or self.channel):
            raise

        params = {'token': self.token,
                  'channel': channel}

        return requests.post(self.api_url, data=msg, params=params)

class IncomingWebHooks(object):
    def __init__(self, token, team, channel=None):
        self.token = token
        self.channel = channel
        self.api_url = 'https://%s.slack.com/services/hooks/incoming-webhook' % team

    def send(self, msg, channel=None, attachments=None):
        if not channel and self.channel:
            channel = self.channel
        elif not (channel or self.channel):
            raise

        payload = {'text': msg,
                   'channel': channel,
                   'username': 'PinkoiBot',
                   'icon_emoji': ':mypinkoi:'}
        if attachments:
            payload.update({'attachments': attachments})

        data = {'payload': json.dumps(payload)}
        params = {'token': self.token}
        return requests.post(self.api_url, data=data, params=params)

    def send_with_attachments(self, msg, channel=None, attachments=None):
        return self.send(msg, channel, attachments)

    @staticmethod
    def render_attachments(fallback, fields_title, fields_value, text=None,
            pretext=None, color="#5060ef", fields_short=False):
        result = {'fallback': fallback,
                  'color': color,
                  'fields': [
                          {'title': fields_title,
                           'value': fields_value,
                           'fields_short': fields_short}
                          ]
                 }
        result.update({"text": text}) if text else None
        result.update({"pretext": pretext}) if pretext else None
        return result

if __name__ == '__main__':
    bot = IncomingWebHooks('...', '...')
    result = bot.send('test', '@...')
    print result.content
