// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ContactComponent() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<section class=\"py-5 py-md-7 bg-light\" id=\"contact\"><div class=\"binary-bg\"></div><div class=\"container\"><h2 class=\"section-heading text-center\">Contact Me</h2><div class=\"container\"><div class=\"mb-4 mb-lg-0\"><div class=\"d-flex flex-column justify-content-center align-items-center border-0 shadow-sm p-4\"><div class=\"container d-flex flex-column align-items-center\"><div class=\"contact-info mb-3\"><i class=\"fas fa-map-marker-alt\"></i> <span>Berlin, Germany</span></div><div class=\"contact-info mb-3\"><i class=\"fas fa-envelope\"></i> <span>alex.gorbunov.de@gmail.com</span></div><div class=\"contact-info mb-3\"><i class=\"fa-brands fa-telegram\"></i> <span>GorbunovAlexander</span></div><div class=\"contact-info mb-3\"><i class=\"fa-brands fa-upwork\"></i> <a href=\"https://www.upwork.com/freelancers/~016c5c40b491593844\" target=\"_blank\"><span>Upwork Top Rated</span></a></div></div><div class=\"social-links d-flex row mt-4\"><a href=\"https://github.com/GorbunovAlex\" target=\"_blank\" class=\"me-2\"><i class=\"fab fa-github\"></i></a> <a href=\"https://www.linkedin.com/in/gorbunovalex/\" target=\"_blank\" class=\"me-2\"><i class=\"fab fa-linkedin\"></i></a> <a href=\"https://x.com/al_gorbunov\" target=\"_blank\" class=\"me-2\"><i class=\"fab fa-twitter\"></i></a></div></div></div></div></div></section><style>\n      @media (max-width: 576px) {\n        .contact-info {\n          display: flex;\n          flex-direction: column;\n          align-items: center;\n          justify-content: center;\n        }\n      }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
