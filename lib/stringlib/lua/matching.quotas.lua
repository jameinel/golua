local quota = require 'quota'

-- Tests for string.find
do
    -- string.find in plain mode consumes cpu proportional to search string
    print(quota.rcall(1000, 0, string.find, ("straw"):rep(20).."needle", "needle", 1, true))
    --> =true	101	106

    print(quota.rcall(1000, 0, string.find, ("straw"):rep(200).."needle", "needle", 1, true))
    --> =false

    -- string.find in pattern mode consumes cpu proportional to the amount of
    -- searching
    print(quota.rcall(10000, 0, string.find, ("a"):rep(50), ".-b"))
    --> =true	nil

    print(quota.rcall(10000, 0, string.find, ("a"):rep(500), ".-b"))
    --> =false

    -- captures consumes memory
    print(quota.rcall(0, 1000, string.find, "abbbbbbbbbbc", "(b+)"))
    --> =true	2	11	bbbbbbbbbb

    print(quota.rcall(0, 3000, string.find, "a"..("b"):rep(1000).."c", "(((b+)))"))
    --> =false
end

-- Tests for string.match
do
    -- string.match consumes cpu proportional to the amount of searching
    print(quota.rcall(10000, 0, string.match, ("a"):rep(50), ".-b"))
    --> =true	nil

    print(quota.rcall(10000, 0, string.match, ("a"):rep(500), ".-b"))
    --> =false

    -- captures consumes memory
    print(quota.rcall(0, 1000, string.match, "abbbbbbbbbbc", "(b+)"))
    --> =true	bbbbbbbbbb

    print(quota.rcall(0, 3000, string.match, "a"..("b"):rep(1000).."c", "(((b+)))"))
    --> =false
end

-- Tests for string.gmatch
do
    local wc
    local function countwords(s)
        wc = 0
        for w in string.gmatch(s, "%w+") do
            wc = wc + 1
        end
    end

    -- every match returned consumes cpu
    print(quota.rcall(1000, 0, countwords, ("hello"):rep(10, " ")))
    --> =true
    print(wc)
    --> =10

    print(quota.rcall(1000, 0, countwords, ("hello"):rep(1000, " ")))
    --> =false
    print(wc > 10 and wc < 200)
    --> =true

    -- every match returned consumes memory
    print(quota.rcall(0, 1000, countwords, ("hello"):rep(10, " ")))
    --> =true
    print(wc)
    --> =10

    print(quota.rcall(0, 1000, countwords, ("hello"):rep(1000, " ")))
    --> =false
    print(wc > 10 and wc < 200)
    --> =true
end
