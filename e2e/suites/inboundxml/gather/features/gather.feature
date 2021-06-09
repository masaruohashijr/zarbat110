Feature: gather    
  In order to input digits and a speech that the call should listen for
  As an end user
  I want to gather the input digits or the speech said.

  Scenario: Gather DTMF

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberA" configured to send digits "123" #-> configuredToSendDigits(numberA string, digits string)
      When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberB" should get digits "123" #-> shouldGetDigits(numberB string, digits string)

  Scenario: Gather Speech

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberA" configured to say "What a wonderful world" #-> configuredToSay(numberB string, message string)
      When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberB" should get speech "What a wonderful world" #-> shouldGetSpeech(numberA string, text string)  
