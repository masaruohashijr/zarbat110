Feature: number    
  In order to try a sequence of numbers until connect
  As an end user
  I want to call to a Number
  And this current call, in case the previous number does not connect, trys to connect to a next number in sequence.

  Scenario: Call multiple numbers in a sequence

    Given I have my "Account"
    When I add to sequence "+777"
    And I add to sequence "+888"
    And I add to sequence "+999"
    Then "+777" did not connect
    And "+888" did not connect
    And "+999" connects