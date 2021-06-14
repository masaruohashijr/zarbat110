Feature: ping    
  In order to send data from a current call to a Webhook
  As an end user
  I want to call to a Number and ping a URL.

  Scenario: Ping a URL
    Given my test setup runs #-> myTestSetupRuns() 
    And "NumberB" configured to ping "webhook" #-> configuredToPing(numberB string, webhook string)
    When I make a call from "NumberA" to "NumberB" #-> iMakeACallFromTo(numberA string, numberB string)
    Then "webhook" should get request method "GET" #-> shouldGetRequestMethod(webhook string, method string)