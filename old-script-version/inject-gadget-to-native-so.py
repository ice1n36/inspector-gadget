#!/usr/bin/env python3

import lief

libnative = lief.parse("output/com.riotgames.league.teamfighttactics/lib/arm64-v8a/libRiotGamesApi.so")
libnative.add_library("libfridagadget.so") # Injection!
libnative.write("output/com.riotgames.league.teamfighttactics/lib/arm64-v8a/libRiotGamesApi.so")
