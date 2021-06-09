Feature: hangup    
  In order to hangup a call after a certain time in seconds
  As an end user
  I want to have my call automatically off

  Scenario: Hangup a call

    Given my test setup runs #-> myTestSetupRuns()
      And "NumberB" configured to hangup after 3 seconds #-> configuredToHangupAfterSeconds(numberB string, seconds int)
      When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
      Then "NumberA" should get last call duration equals to 3 #-> shouldGetLastCallDurationEqualsTo(numberA string, timeInSeconds int)