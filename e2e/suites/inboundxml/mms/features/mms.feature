Feature: mms    
  In order to my sent MMS messages can be seem by my destination numbers
  And getting a receiving confirmation from them
  As an end user
  I want to send an MMS message

  Scenario: Send an MMS

    Given I have my "Account"
    When I set an MMS message with "test.jpg" and "And I think to myself"
    And I send the message from "+888" to "+999" 
    Then the "+999" should receive an MMS 
    And message contains media "test.jpg" 
    And message contains text "And I think to myself"