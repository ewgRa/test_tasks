describe('All in one', () => {
  it('Generate short url and visit it', () => {

    cy.visit('http://localhost:8080');

    cy.get('#longUrlInput')
      .type('http://long-url.com')
      .should('have.value', 'http://long-url.com');

    cy.get('#makeButton').click();

    cy.get('#successResult', { timeout: 5000 }).should('be.visible');
    cy.get('#longUrl').contains('http://long-url.com');

    cy.get('#shortUrl').invoke("text").then((text) => {
      cy.request({
        url: text,
        followRedirect: false
      })
        .then((resp) => {
          expect(resp.status).to.eq(302)
          expect(resp.redirectedToUrl).to.eq('http://long-url.com/')
        });
    })
  });
})