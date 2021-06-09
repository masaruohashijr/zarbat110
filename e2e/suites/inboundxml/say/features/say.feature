Feature: say    
  In order to read text to the Number B (called) using a text-to-speech engine   
  As an end user
  I want that Number A (caller) listen the speech set to be read on Number B.

  Scenario: Say something

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberB" configured to say "I think to myself" #-> configuredToSay(numberA string, message string)
      And "NumberA" configured to gather speech #-> configuredToGatherSpeech(numberB string)
      When I make a call from "NumberA" to "NumbeB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberA" should get speech "I think to myself" #-> shouldGetSpeech(numberA string, speech string)    
