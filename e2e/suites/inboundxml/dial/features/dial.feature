Feature: dial    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call

Scenario: Dial
  Given my test setup runs #-> myTestSetupRuns()
    And "NumberA" configured to say "Hello World" #-> configuredToSay(numberA string, message string)
    And append to "NumberA" config dial "NumberC" #-> appendToConfigDial(numberA string, numberC string)
    When I make a call from "NumberB" to "NumberA" #-> iMakeACallFromTo(numberB string, numberA string)
    Then "NumberC" should get the incoming call from "NumberA" #-> shouldGetTheIncomingCallFrom(numberC string, numberA string)
    And "NumberA" should get the incoming call from "NumberB" #-> shouldGetTheIncomingCallFrom(numberA string, numberB string)