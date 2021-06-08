Feature: ping    
  In order to send data from a current call to a Webhook
  As an end user
  I want to call to a Number and ping a URL.

  Scenario: Ping an URL

    Given I have my "Account"
    And I set "+999" to ping "http://webhook"
    When I make a call from "+888" to "+999"
    Then I get voice request parameters