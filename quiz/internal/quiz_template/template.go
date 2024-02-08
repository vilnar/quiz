package quiz_template

import (
	"path"
	"quiz/internal/common"
)

func GetFilesForParseReport(main, header, footer string) []string {
	return []string{
		main,
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "kotenov_5_57_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "first_ptsd_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "nps_prognoz_2_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "hads_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "ies_r_5_54_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "minimult_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "iso_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "stai_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "eysenck_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "lnp_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "qsr_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "dfp_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "adaptability_200_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "occupational_burnout_result_content.html"),
		header,
		footer,
		// TODO: find in template quiz_switch_todo
	}
}
