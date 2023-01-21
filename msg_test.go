package main

import (
	"fmt"
	"net/mail"
	"testing"
)

func TestFormat(t *testing.T) {
	x := `  {
    "id": "109722145263655254",
    "created_at": "2023-01-20T15:04:30.067Z",
    "in_reply_to_id": null,
    "in_reply_to_account_id": null,
    "sensitive": false,
    "spoiler_text": "",
    "visibility": "public",
    "language": null,
    "uri": "https://social.treehouse.systems/users/alyssa/statuses/109722145263655254/activity",
    "url": "https://social.treehouse.systems/users/alyssa/statuses/109722145263655254/activity",
    "replies_count": 0,
    "reblogs_count": 0,
    "favourites_count": 0,
    "edited_at": null,
    "favourited": false,
    "reblogged": false,
    "muted": false,
    "bookmarked": false,
    "local_only": false,
    "content": "",
    "filtered": [],
    "reblog": {
      "id": "109721809783205964",
      "created_at": "2023-01-20T13:39:13.000Z",
      "in_reply_to_id": null,
      "in_reply_to_account_id": null,
      "sensitive": false,
      "spoiler_text": "",
      "visibility": "public",
      "language": "en",
      "uri": "https://social.hackerspace.pl/users/q3k/statuses/109721809926740193",
      "url": "https://social.hackerspace.pl/@q3k/109721809926740193",
      "replies_count": 9,
      "reblogs_count": 21,
      "favourites_count": 4,
      "edited_at": null,
      "favourited": false,
      "reblogged": false,
      "muted": false,
      "bookmarked": false,
      "local_only": null,
      "content": "<p>obligatory neofetch, because apparently people are into that</p>",
      "filtered": [],
      "reblog": null,
      "account": {
        "id": "109360173917975335",
        "username": "q3k",
        "acct": "q3k@hackerspace.pl",
        "display_name": "q3k :blobcatcoffee:",
        "locked": false,
        "bot": false,
        "discoverable": false,
        "group": false,
        "created_at": "2022-11-17T00:00:00.000Z",
        "note": "<p>Documenting the hyperfocus episodes of a soul stuck between hardware and software. </p><p>THIS CONTENT IS PROVIDED BY THE AUTHOR \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.</p><p><a href=\"https://social.hackerspace.pl/tags/nobot\" class=\"mention hashtag\" rel=\"nofollow noopener noreferrer\" target=\"_blank\">#<span>nobot</span></a></p><p>(Old account: <a href=\"https://0x3c.pl/@q3k\" rel=\"nofollow noopener noreferrer\" target=\"_blank\"><span class=\"invisible\">https://</span><span class=\"\">0x3c.pl/@q3k</span><span class=\"invisible\"></span></a>)</p>",
        "url": "https://social.hackerspace.pl/@q3k",
        "avatar": "https://cache.treehouse.systems/cache/accounts/avatars/109/360/173/917/975/335/original/a1e0ab99099b325a.png",
        "avatar_static": "https://cache.treehouse.systems/cache/accounts/avatars/109/360/173/917/975/335/original/a1e0ab99099b325a.png",
        "header": "https://cache.treehouse.systems/cache/accounts/headers/109/360/173/917/975/335/original/08817afa04f60c9a.jpeg",
        "header_static": "https://cache.treehouse.systems/cache/accounts/headers/109/360/173/917/975/335/original/08817afa04f60c9a.jpeg",
        "followers_count": 371,
        "following_count": 146,
        "statuses_count": 115,
        "last_status_at": "2023-01-21",
        "emojis": [
          {
            "shortcode": "blobcatcoffee",
            "url": "https://cache.treehouse.systems/cache/custom_emojis/images/000/044/523/original/e1783ce6bf48667f.png",
            "static_url": "https://cache.treehouse.systems/cache/custom_emojis/images/000/044/523/static/e1783ce6bf48667f.png",
            "visible_in_picker": true
          }
        ],
        "fields": [
          {
            "name": "Location",
            "value": "Munich, Germany",
            "verified_at": null
          },
          {
            "name": "Pronouns",
            "value": "he/him",
            "verified_at": null
          },
          {
            "name": "Contact",
            "value": "<a href=\"https://q3k.org/\" rel=\"nofollow noopener noreferrer\" target=\"_blank\"><span class=\"invisible\">https://</span><span class=\"\">q3k.org/</span><span class=\"invisible\"></span></a>",
            "verified_at": "2023-01-18T08:58:55.548+00:00"
          }
        ]
      },
      "media_attachments": [
        {
          "id": "109721809694512674",
          "type": "image",
          "url": "https://cache.treehouse.systems/cache/media_attachments/files/109/721/809/694/512/674/original/136141c61d5aac10.png",
          "preview_url": "https://cache.treehouse.systems/cache/media_attachments/files/109/721/809/694/512/674/small/136141c61d5aac10.png",
          "remote_url": "https://object.ceph-waw3.hswaw.net/mastodon-prod/media_attachments/files/109/721/808/577/636/905/original/b83d41ac0734968e.png",
          "preview_remote_url": null,
          "text_url": null,
          "meta": {
            "focus": {
              "x": 0,
              "y": 0
            },
            "original": {
              "width": 776,
              "height": 401,
              "size": "776x401",
              "aspect": 1.9351620947630923
            },
            "small": {
              "width": 668,
              "height": 345,
              "size": "668x345",
              "aspect": 1.936231884057971
            }
          },
          "description": "root@nano5g\n-----------\nOS: Alpine Linux v3.17 armv6l\nHost: Apple iPod Nano 5G\nKernel: 6.2.0-rc3-00024-ge2e3252e9e4e-dirty\nUptime: 40 mins\nPackages: 67 (apk)\nShell: sh\nTerminal: /dev/pts/0\nCPU: Generic DT based system (1)\nMemory: 29MiB / 51MiB",
          "blurhash": "U55F5*t:IV$_s;fkfhe:DioH%LMyx]ahRRtQ"
        }
      ],
      "mentions": [],
      "tags": [],
      "emojis": [],
      "card": null,
      "poll": null,
      "quote": null
    },
    "account": {
      "id": "109311456944646396",
      "username": "alyssa",
      "acct": "alyssa",
      "display_name": "Alyssa Rosenzweig 💜",
      "locked": false,
      "bot": false,
      "discoverable": true,
      "group": false,
      "created_at": "2022-11-09T00:00:00.000Z",
      "note": "<p>Linux hacker tinkering with graphics drivers</p>",
      "url": "https://social.treehouse.systems/@alyssa",
      "avatar": "https://cache.treehouse.systems/accounts/avatars/109/311/456/944/646/396/original/554f453a50fcf60d.png",
      "avatar_static": "https://cache.treehouse.systems/accounts/avatars/109/311/456/944/646/396/original/554f453a50fcf60d.png",
      "header": "https://cache.treehouse.systems/accounts/headers/109/311/456/944/646/396/original/495cc6d3c6f89d40.png",
      "header_static": "https://cache.treehouse.systems/accounts/headers/109/311/456/944/646/396/original/495cc6d3c6f89d40.png",
      "followers_count": 6737,
      "following_count": 19,
      "statuses_count": 211,
      "last_status_at": "2023-01-21",
      "noindex": false,
      "emojis": [],
      "fields": [
        {
          "name": "Cybre home",
          "value": "<a href=\"https://rosenzweig.io/\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\"><span class=\"invisible\">https://</span><span class=\"\">rosenzweig.io/</span><span class=\"invisible\"></span></a>",
          "verified_at": "2022-11-09T02:35:52.825+00:00"
        },
        {
          "name": "Physical home",
          "value": "North America",
          "verified_at": null
        }
      ]
    },
    "media_attachments": [],
    "mentions": [],
    "tags": [],
    "emojis": [],
    "card": null,
    "poll": null,
    "quote": null
  }
`
	out := formatMessage(rawMessage{id: 123, content: x},
		&mail.Address{Name: "Ro Bot", Address: "ro@b.ot"},
		&mail.Address{Name: "Hello", Address: "wo@r.ld"})
	fmt.Printf("%s\n", out)
}
