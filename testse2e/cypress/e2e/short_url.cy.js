describe('short url page', () => {
  beforeEach(() => {
    cy.visit('http://localhost:8081/app/')
  })

  it('returns error if url empty', () => {
    cy.get('.formContainer > md-outlined-button:nth-child(2)').click();
    cy.get('.formContainer > md-outlined-text-field:nth-child(1)')
      .should('have.attr', 'error-text', 'No Link provided');
  })

  it('return short link', () => {
    cy.get('.formContainer > md-outlined-text-field:nth-child(1)')
      .shadow()
      .find('input')
      .click()
      .type('https://google.com');

    cy.get('.formContainer > md-outlined-button:nth-child(2)').click();

    cy.get('#shorten_url').should('have.attr', 'value')
  })
})