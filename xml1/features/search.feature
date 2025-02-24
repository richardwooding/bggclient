Feature: Search for boardgames

  Scenario: Search for a specific boardgame
    Given the API is initialized with a valid base URL and HTTP client
    When I search for "Catan"
    Then I should receive a list of boardgames
    And the list should contain a boardgame with the name "Catan"

  Scenario: Search for a non-existent boardgame
    Given the API is initialized with a valid base URL and HTTP client
    When I search for "NonExistentGame"
    Then I should receive an empty list of boardgames

  Scenario: Search with an empty string
    Given the API is initialized with a valid base URL and HTTP client
    When I search for an empty string
    Then I should receive an empty list of boardgames