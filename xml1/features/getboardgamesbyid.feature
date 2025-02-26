Feature: Get boardgames by ID

  Scenario: Get a specific boardgame by ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "12345"
    Then I should receive a single boardgame
    And the boardgame should have the ID "12345"

  Scenario: Get multiple boardgames by IDs
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgames with IDs
      | 12345 |
      | 67890 |
    Then I should receive a list of boardgames
    And the list should contain a boardgame with the IDs
      | 12345 |
      | 67890 |

  Scenario: Get mure then 20 boardgames by IDs
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgames with IDs
      | 12345 |
      | 67890 |
      | 67891 |
      | 67892 |
      | 67893 |
      | 67894 |
      | 67895 |
      | 67896 |
      | 67897 |
      | 67898 |
      | 67899 |
      | 67900 |
      | 67901 |
      | 67902 |
      | 67903 |
      | 67904 |
      | 67905 |
      | 67906 |
      | 67907 |
      | 67908 |
      | 67909 |
      | 67910 |
      | 67911 |
    Then I should receive an error message
    And the error message should indicate that the number of IDs is invalid


  Scenario: Get a non-existent boardgame by ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "00000"
    Then I should receive an error message
    And the error message should indicate that the boardgame was not found

  Scenario: Get boardgames with an empty ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with an empty ID
    Then I should receive an error message
    And the error message should indicate that the ID is invalid

  Scenario: Get boardgames with comments included
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "12345" with comments included
    Then I should receive a single boardgame
    And the boardgame should have the ID "12345"
    And the boardgame should have comments

  Scenario: Get boardgames with stats included
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "12345" with stats included
    Then I should receive a single boardgame
    And the boardgame should have the ID "12345"
    And the boardgame should have stats

  Scenario: Get boardgames with historical data included
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "12345" with historical data included from 2009-01-01 to 2009-03-17
    Then I should receive a single boardgame
    And the boardgame should have the ID "12345"
    # And the boardgame should have historical data
