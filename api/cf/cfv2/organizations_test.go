package cfv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/client"
	bluemixHttp "github.com/IBM-Bluemix/bluemix-go/http"
	"github.com/IBM-Bluemix/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Organizations", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("Create", func() {
		Context("When creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/organizations"),
						ghttp.VerifyBody([]byte(`{"name":"test-org"}`)),
						ghttp.RespondWith(http.StatusCreated, `{
							 "metadata": {
							    "guid": "a87f1cde-2070-4fb1-b23e-09109d9eaa93",
							    "url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93",
							    "created_at": "2016-04-16T01:23:42Z",
							    "updated_at": null
							  },
							  "entity": {
							    "name": "test-org",
							    "billing_enabled": false,
							    "quota_definition_guid": "813ca4a3-e83f-463e-a034-3a7ed7ba6280",
							    "status": "active",
							    "quota_definition_url": "/v2/quota_definitions/813ca4a3-e83f-463e-a034-3a7ed7ba6280",
							    "spaces_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/spaces",
							    "domains_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/domains",
							    "private_domains_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/private_domains",
							    "users_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/users",
							    "managers_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/managers",
							    "billing_managers_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/billing_managers",
							    "auditors_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/auditors",
							    "app_events_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/app_events",
							    "space_quota_definitions_url": "/v2/organizations/a87f1cde-2070-4fb1-b23e-09109d9eaa93/space_quota_definitions"
							  }							
						}`),
					),
				)
			})

			It("Should create Organization", func() {
				err := newOrganizations(server.URL()).Create("test-org")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/organizations"),
						ghttp.VerifyBody([]byte(`{"name":"test-org"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create`),
					),
				)
			})

			It("Should return error when created", func() {
				err := newOrganizations(server.URL()).Create("test-org")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("FindByName", func() {
		Context("When there is match", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations", "q=name:foo"),
						ghttp.RespondWith(http.StatusOK, `{
							  "total_results": 1,
							  "total_pages": 1,
							  "prev_url": null,
							  "next_url": null,
							  "resources": [
							    {
							      "metadata": {
							        "guid": "a695a906-e225-428a-9238-b1d01b96017f",
							        "url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f",
							        "created_at": "2016-04-16T01:23:41Z",
							        "updated_at": null
							      },
							      "entity": {
							        "name": "foo",
							        "billing_enabled": false,
							        "quota_definition_guid": "255876b7-4f1c-48d9-ad7c-2d6cd590781e",
							        "status": "active",
							        "quota_definition_url": "/v2/quota_definitions/255876b7-4f1c-48d9-ad7c-2d6cd590781e",
							        "spaces_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/spaces",
							        "domains_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/domains",
							        "private_domains_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/private_domains",
							        "users_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/users",
							        "managers_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/managers",
							        "billing_managers_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/billing_managers",
							        "auditors_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/auditors",
							        "app_events_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/app_events",
							        "space_quota_definitions_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/space_quota_definitions"
							      }
							    }
							  ]
						}`),
					),
				)
			})

			It("Should return the match", func() {
				org, err := newOrganizations(server.URL()).FindByName("foo")
				Expect(err).NotTo(HaveOccurred())
				Expect(org).NotTo(BeNil())
				Expect(org.GUID).To(Equal("a695a906-e225-428a-9238-b1d01b96017f"))
				Expect(org.Name).To(Equal("foo"))
				Expect(org.BillingEnabled).To(BeFalse())
			})
		})

		Context("When Organizations FindByName is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations", "q=name:foo"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to find by name`),
					),
				)
			})

			It("Should return error when Organization is found by name", func() {
				org, err := newOrganizations(server.URL()).FindByName("foo")
				Expect(err).To(HaveOccurred())
				Expect(org).Should(BeNil())
			})
		})
	})

	Describe("List", func() {
		Context("When there is one Organization", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations"),
						ghttp.RespondWith(http.StatusOK, `{
							  "total_results": 1,
							  "total_pages": 1,
							  "prev_url": null,
							  "next_url": null,
							  "resources": [
							    {
							      "metadata": {
							        "guid": "a695a906-e225-428a-9238-b1d01b96017f",
							        "url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f",
							        "created_at": "2016-04-16T01:23:41Z",
							        "updated_at": null
							      },
							      "entity": {
							        "name": "test-org-name",
							        "billing_enabled": false,
							        "quota_definition_guid": "255876b7-4f1c-48d9-ad7c-2d6cd590781e",
							        "status": "active",
							        "quota_definition_url": "/v2/quota_definitions/255876b7-4f1c-48d9-ad7c-2d6cd590781e",
							        "spaces_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/spaces",
							        "domains_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/domains",
							        "private_domains_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/private_domains",
							        "users_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/users",
							        "managers_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/managers",
							        "billing_managers_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/billing_managers",
							        "auditors_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/auditors",
							        "app_events_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/app_events",
							        "space_quota_definitions_url": "/v2/organizations/a695a906-e225-428a-9238-b1d01b96017f/space_quota_definitions"
							      }
							    }
							  ]
						}`),
					),
				)
			})

			It("Should return all Organizations", func() {
				orgs, err := newOrganizations(server.URL()).List()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(orgs).Should(HaveLen(1))
				org := orgs[0]
				Expect(org.GUID).To(Equal("a695a906-e225-428a-9238-b1d01b96017f"))
				Expect(org.Name).To(Equal("test-org-name"))
				Expect(org.BillingEnabled).To(BeFalse())
			})
		})

		Context("When Organizations List is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get Organizations`),
					),
				)
			})

			It("Should return error when Organization List is failed", func() {
				orgs, err := newOrganizations(server.URL()).List()
				Expect(err).To(HaveOccurred())
				Expect(orgs).Should(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When Organization deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/organizations/827ec655-c2ed-4577-b226-82271fe471d7"),
						ghttp.RespondWith(http.StatusNoContent, `
						{}`),
					),
				)
			})

			It("Should delete Organization", func() {
				err := newOrganizations(server.URL()).Delete("827ec655-c2ed-4577-b226-82271fe471d7", true)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When Organization Delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/organizations/827ec655-c2ed-4577-b226-82271fe471d7"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete`),
					),
				)
			})

			It("Should return error when Organization is deleted", func() {
				err := newOrganizations(server.URL()).Delete("827ec655-c2ed-4577-b226-82271fe471d7", true)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update", func() {
		Context("When update is succeesful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c"),
						ghttp.VerifyBody([]byte(`{"name":"new-org-name"}`)),
						ghttp.RespondWith(http.StatusCreated, `{
						  "metadata": {
						    "guid": "007c547f-9d6e-4d75-bb03-d9584e7bc62c",
						    "url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c",
						    "created_at": "2016-04-16T01:23:42Z",
						    "updated_at": "2016-04-16T01:23:42Z"
						  },
						  "entity": {
						    "name": "new-org-name",
						    "billing_enabled": false,
						    "quota_definition_guid": "838224b1-c841-4b96-90f2-f181bb5fffa5",
						    "status": "active",
						    "quota_definition_url": "/v2/quota_definitions/838224b1-c841-4b96-90f2-f181bb5fffa5",
						    "spaces_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/spaces",
						    "domains_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/domains",
						    "private_domains_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/private_domains",
						    "users_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/users",
						    "managers_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/managers",
						    "billing_managers_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/billing_managers",
						    "auditors_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/auditors",
						    "app_events_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/app_events",
						    "space_quota_definitions_url": "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c/space_quota_definitions"
						  }						
						}`),
					),
				)
			})

			It("Should update Organization", func() {
				err := newOrganizations(server.URL()).Update("007c547f-9d6e-4d75-bb03-d9584e7bc62c", "new-org-name")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When Organization Update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/organizations/007c547f-9d6e-4d75-bb03-d9584e7bc62c"),
						ghttp.VerifyBody([]byte(`{"name":"new-org-name"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update`),
					),
				)
			})

			It("Should return error when updated", func() {
				err := newOrganizations(server.URL()).Update("007c547f-9d6e-4d75-bb03-d9584e7bc62c", "new-org-name")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func newOrganizations(url string) Organizations {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	cfClient := client.Client{
		Config:           conf,
		ServiceName:      bluemix.CfService,
		HandlePagination: Paginate,
	}

	return newOrganizationAPI(&cfClient)
}
