package builds

import (
	"fmt"

	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	exutil "github.com/openshift/origin/test/extended/util"
)

var _ = g.Describe("builds: parallel: build revision", func() {
	defer g.GinkgoRecover()
	var (
		buildFixture = exutil.FixturePath("..", "extended", "fixtures", "test-build-revision.json")
		oc           = exutil.NewCLI("cli-build-revision", exutil.KubeConfigPath())
	)

	g.JustBeforeEach(func() {
		g.By("waiting for builder service account")
		err := exutil.WaitForBuilderAccount(oc.KubeREST().ServiceAccounts(oc.Namespace()))
		o.Expect(err).NotTo(o.HaveOccurred())
		oc.Run("create").Args("-f", buildFixture).Execute()
	})

	g.Describe("started build", func() {
		g.It("should contain source revision information", func() {
			g.By("starting the build with --wait flag")
			out, err := oc.Run("start-build").Args("sample-build", "--wait").Output()
			o.Expect(err).NotTo(o.HaveOccurred())

			g.By(fmt.Sprintf("verifying the build %q status", out))
			build, err := oc.REST().Builds(oc.Namespace()).Get(out)
			o.Expect(err).NotTo(o.HaveOccurred())
			o.Expect(build.Spec.Revision).NotTo(o.BeNil())
			o.Expect(build.Spec.Revision.Git).NotTo(o.BeNil())
			o.Expect(build.Spec.Revision.Git.Commit).NotTo(o.BeEmpty())
			o.Expect(build.Spec.Revision.Git.Author.Name).NotTo(o.BeEmpty())
			o.Expect(build.Spec.Revision.Git.Committer.Name).NotTo(o.BeEmpty())
			o.Expect(build.Spec.Revision.Git.Message).NotTo(o.BeEmpty())
		})
	})
})
