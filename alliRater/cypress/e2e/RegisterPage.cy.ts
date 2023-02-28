describe('Register Page Test', () => {
  it('Visits the Homepage', () => {
    cy.visit('http://localhost:4200/')

    cy.contains('Sign Up').click()

    // Should be on a new URL which
    // includes '/register'
    cy.url().should('include', '/register')
  })
})