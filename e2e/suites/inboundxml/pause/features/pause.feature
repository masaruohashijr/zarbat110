Feature: pause    
  In order to pause a sequence of phrases
  As an end user
  I want to call to a Number and say some words or phrases pausing some seconds between each sentence.

  Scenario: Pause a sequence of sentences

    Given my test setup runs #-> myTestSetupRuns()
    And "NumberB" configured to pause 5 seconds #-> configuredToPauseSeconds(numberB string, timeInSeconds int)
    And append To "NumberB" config hangup #-> appendToConfigHangup(numberB string)
    When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
    Then "NumberA" should get last call duration equals to 5 #-> shouldGetLastCallDurationEqualsTo(numberB string, timeInSeconds int)
