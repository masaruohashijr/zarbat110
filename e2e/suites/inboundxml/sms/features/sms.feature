Feature: sms    
  In order to my sent SMS messages can be seem by my destination numbers
  And getting a receiving confirmation from them
  As an end user
  I want to send an MMS message

  Scenario: Send an SMS

    Given I have my "Account"
    When I set an SMS message with "What a wonderful world"
    And I send the message from "+888" to "+999" 
    Then the "+999" should receive an SMS
    And message contains text "What a wonderful world"