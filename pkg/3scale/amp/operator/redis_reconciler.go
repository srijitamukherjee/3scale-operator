package operator

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	appsv1 "github.com/openshift/api/apps/v1"
	imagev1 "github.com/openshift/api/image/v1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type RedisReconciler struct {
	BaseAPIManagerLogicReconciler
}

// blank assignment to verify that BaseReconciler implements reconcile.Reconciler
var _ LogicReconciler = &RedisReconciler{}

func NewRedisReconciler(baseAPIManagerLogicReconciler BaseAPIManagerLogicReconciler) RedisReconciler {
	return RedisReconciler{
		BaseAPIManagerLogicReconciler: baseAPIManagerLogicReconciler,
	}
}

func (r *RedisReconciler) Reconcile() (reconcile.Result, error) {
	if r.apiManager.Spec.HighAvailability != nil && r.apiManager.Spec.HighAvailability.Enabled {
		return reconcile.Result{}, nil
	}

	redis, err := r.redis()
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileBackendDeploymentConfig(redis.BackendDeploymentConfig())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileBackendService(redis.BackendService())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileBackendConfigMap(redis.BackendConfigMap())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileBackendPVC(redis.BackendPVC())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileBackendImageStream(redis.BackendImageStream())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileSystemDeploymentConfig(redis.SystemDeploymentConfig())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileSystemPVC(redis.SystemPVC())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileSystemImageStream(redis.SystemImageStream())
	if err != nil {
		return reconcile.Result{}, err
	}

	err = r.reconcileSystemService(redis.SystemService())
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *RedisReconciler) redis() (*component.Redis, error) {
	optsProvider := OperatorRedisOptionsProvider{APIManagerSpec: &r.apiManager.Spec}
	opts, err := optsProvider.GetRedisOptions()
	if err != nil {
		return nil, err
	}
	return component.NewRedis(opts), nil
}

func (r *RedisReconciler) reconcileDeploymentConfig(desiredDeploymentConfig *appsv1.DeploymentConfig) error {
	err := r.InitializeAsAPIManagerObject(desiredDeploymentConfig)
	if err != nil {
		return err
	}

	return r.deploymentConfigReconciler.Reconcile(desiredDeploymentConfig)
}

func (r *RedisReconciler) reconcileConfigMap(desiredConfigMap *v1.ConfigMap) error {
	err := r.InitializeAsAPIManagerObject(desiredConfigMap)
	if err != nil {
		return err
	}

	return r.configMapReconciler.Reconcile(desiredConfigMap)
}

func (r *RedisReconciler) reconcilePersistentVolumeClaim(desiredPVC *v1.PersistentVolumeClaim) error {
	err := r.InitializeAsAPIManagerObject(desiredPVC)
	if err != nil {
		return err
	}

	return r.persistentVolumeClaimReconciler.Reconcile(desiredPVC)
}

func (r *RedisReconciler) reconcileService(desiredService *v1.Service) error {
	err := r.InitializeAsAPIManagerObject(desiredService)
	if err != nil {
		return err
	}
	return r.serviceReconciler.Reconcile(desiredService)
}

func (r *RedisReconciler) reconcileImageStream(desiredImageStream *imagev1.ImageStream) error {
	err := r.InitializeAsAPIManagerObject(desiredImageStream)
	if err != nil {
		return err
	}

	return r.imagestreamReconciler.Reconcile(desiredImageStream)
}

func (r *RedisReconciler) reconcileBackendDeploymentConfig(desiredDeploymentConfig *appsv1.DeploymentConfig) error {
	return r.reconcileDeploymentConfig(desiredDeploymentConfig)
}

func (r *RedisReconciler) reconcileBackendService(desiredService *v1.Service) error {
	return r.reconcileService(desiredService)
}

func (r *RedisReconciler) reconcileBackendConfigMap(desiredConfigMap *v1.ConfigMap) error {
	return r.reconcileConfigMap(desiredConfigMap)
}

func (r *RedisReconciler) reconcileBackendPVC(desiredPVC *v1.PersistentVolumeClaim) error {
	return r.reconcilePersistentVolumeClaim(desiredPVC)
}

func (r *RedisReconciler) reconcileSystemDeploymentConfig(desiredDeploymentConfig *appsv1.DeploymentConfig) error {
	return r.reconcileDeploymentConfig(desiredDeploymentConfig)
}

func (r *RedisReconciler) reconcileSystemPVC(desiredPVC *v1.PersistentVolumeClaim) error {
	return r.reconcilePersistentVolumeClaim(desiredPVC)
}

func (r *RedisReconciler) reconcileSystemService(desiredService *v1.Service) error {
	return r.reconcileService(desiredService)
}

func (r *RedisReconciler) reconcileBackendImageStream(desiredImageStream *imagev1.ImageStream) error {
	return r.reconcileImageStream(desiredImageStream)
}

func (r *RedisReconciler) reconcileSystemImageStream(desiredImageStream *imagev1.ImageStream) error {
	return r.reconcileImageStream(desiredImageStream)
}
